#!/usr/bin/env perl
# Copyright ©2014 The gonum Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $cblasHeader = "cblas.h";
my $clapackHeader = "clapack.h";
my $LIB = "/usr/lib/";

my $excludeComplex = 0;
my $excludeAtlas = 1;


open(my $clapack, "<", $clapackHeader) or die;
open(my $golapack, ">", "lapack.go") or die;

my %done;

printf $golapack <<EOH;
// Do not manually edit this file. It was created by the genLapack.pl script from ${clapackHeader}.

// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lapack

/*
#cgo CFLAGS: -g -O2 -fPIC -m64 -pthread
#cgo LDFLAGS: -L${LIB} -llapack -lblas
#include "${cblasHeader}"
#include "${clapackHeader}"
*/
import "C"

import (
	"github.com/gonum/blas"
	"unsafe"
)

EOH

$/ = undef;
my $header = <$clapack>;

# horrible munging of text...
$header =~ s/#[^\n\r]*//g;                 # delete cpp lines
$header =~ s/\n +([^\n\r]*)/\n$1/g;        # remove starting space
$header =~ s/(?:\n ?\n)+/\n/g;             # delete empty lines
$header =~ s! ((['"]) (?: \\. | .)*? \2) | # skip quoted strings
             /\* .*? \*/ |                 # delete C comments
             // [^\n\r]*                   # delete C++ comments just in case
             ! $1 || ' '                   # change comments to a single space
             !xseg;    	                   # ignore white space, treat as single line
                                           # evaluate result, repeat globally
$header =~ s/([^;])\n/$1/g;                # join prototypes into single lines
$header =~ s/, +/,/g;
$header =~ s/ +/ /g;
$header =~ s/ +}/}/g;
$header =~ s/\n+//;

$/ = "\n";
my @lines = split ";\n", $header;

our %retConv = (
	"int" => "int ",
	"float" => "float32 ",
	"double" => "float64 ",
	"CBLAS_INDEX" => "int ",
	"void" => ""
);

foreach my $line (@lines) {
	process($line);	
}

close($golapack);
`go fmt .`;

sub process {
	my $line = shift;
	chomp $line;
	processProto($line);
}

sub processProto {
	my $proto = shift;
	my ($func, $paramList) = split /[()]/, $proto;
	(my $ret, $func) = split ' ', $func;
	if ($done{$func} or $excludeComplex && $func =~ m/_[isd]?[zc]/ or $excludeAtlas && $func =~ m/^catlas_/) {
		return
	}
	$done{$func} = 1;
	my $GoRet = $retConv{$ret};
	my $complexType = $func;
	$complexType =~ s/.*_[isd]?([zc]).*/$1/;
	print $golapack "func ".Gofunc($func)."(".processParamToGo($func, $paramList, $complexType).") ".$GoRet."{\n";
	print $golapack "\t";
	if ($ret ne 'void') {
		chop($GoRet);
		print $golapack "return ".$GoRet."(";
	}
	print $golapack "C.$func(".processParamToC($func, $paramList).")";
	if ($ret ne 'void') {
		print $golapack ")";
	}
	print $golapack "\n}\n";
}

sub Gofunc {
	my $fnName = shift;
	my ($pack, $func, $tail) = split '_', $fnName;
	if ($pack eq 'clapack') {
		$pack = "";
	} else {
		$pack = substr $pack, 1;
	}
	return ucfirst $pack . ucfirst $func . ucfirst $tail if $tail;
	return ucfirst $pack . ucfirst $func;
}

sub processParamToGo {
	my $func = shift;
	my $paramList = shift;
	my $complexType = shift;
	my @processed;
	my @params = split ',', $paramList;
	foreach my $param (@params) {
		my @parts = split /[ *]/, $param;
		my $var = lcfirst $parts[scalar @parts - 1];
		$param =~ m/^(?:const )?int \*/ && do {
			push @processed, $var." []int32"; next;
		};
		$param =~ m/^(?:const )?int/ && do {
			push @processed, $var." int"; next;
		};
		$param =~ m/^(?:const )?void/ && do {
			my $type;
			if ($var eq "alpha" || $var eq "beta") {
				$type = " ";
			} else {
				$type = " []";
			}
			if ($complexType eq 'c') {
				push @processed, $var.$type."complex64"; next;
			} elsif ($complexType eq 'z') {
				push @processed, $var.$type."complex128"; next;
			} else {
				die "unexpected complex type for '$func' - '$complexType'";
			}
		};
		$param =~ m/^(?:const )?char \*/ && do {
			push @processed, $var." *byte"; next;
		};
		$param =~ m/^(?:const )?float \*/ && do {
			push @processed, $var." []float32"; next;
		};
		$param =~ m/^(?:const )?double \*/ && do {
			push @processed, $var." []float64"; next;
		};
		$param =~ m/^(?:const )?float/ && do {
			push @processed, $var." float32"; next;
		};
		$param =~ m/^(?:const )?double/ && do {
			push @processed, $var." float64"; next;
		};
		$param =~ m/^const enum/ && do {
			$var eq "order" && do {
				$var = "o";
				push @processed, $var." blas.Order"; next;
			};
			$var =~ /trans/ && do {
				$var =~ s/trans([AB]?)/t$1/;
				push @processed, $var." blas.Transpose"; next;
			};
			$var eq "uplo" && do {
				$var = "ul";
				push @processed, $var." blas.Uplo"; next;
			};
			$var eq "diag" && do {
				$var = "d";
				push @processed, $var." blas.Diag"; next;
			};
			$var eq "side" && do {
				$var = "s";
				push @processed, $var." blas.Side"; next;
			};
		};
	}
	die "missed Go parameters from '$func', '$paramList'" if scalar @processed != scalar @params;
	return join ", ", @processed;
}

sub processParamToC {
	my $func = shift;
	my $paramList = shift;
	my @processed;
	my @params = split ',', $paramList;
	foreach my $param (@params) {
		my @parts = split /[ *]/, $param;
		my $var = lcfirst $parts[scalar @parts - 1];
		$param =~ m/^(?:const )?int \*[a-zA-Z]/ && do {
			push @processed, "(*C.int)(&".$var."[0])"; next;
		};
		$param =~ m/^(?:const )?void \*[a-zA-Z]/ && do {
			my $type;
			if ($var eq "alpha" || $var eq "beta") {
				$type = "";
			} else {
				$type = "[0]";
			}
			push @processed, "unsafe.Pointer(&".$var.$type.")"; next;
		};
		$param =~ m/^(?:const )?char \*[a-zA-Z]/ && do {
			push @processed, "(*C.char)(&".$var.")"; next;
		};
		$param =~ m/^(?:const )?float \*[a-zA-Z]/ && do {
			push @processed, "(*C.float)(&".$var."[0])"; next;
		};
		$param =~ m/^(?:const )?double \*[a-zA-Z]/ && do {
			push @processed, "(*C.double)(&".$var."[0])"; next;
		};
		$param =~ m/^(?:const )?int [a-zA-Z]/ && do {
			push @processed, "C.int(".$var.")"; next;
		};
		$param =~ m/^(?:const )float [a-zA-Z]/ && do {
			push @processed, "C.float(".$var.")"; next;
		};
		$param =~ m/^(?:const )double [a-zA-Z]/ && do {
			push @processed, "C.double(".$var.")"; next;
		};
		$param =~ m/^const enum [a-zA-Z]/ && do {
			$var eq "order" && do {
				$var = "o";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var =~ /trans/ && do {
				$var =~ s/trans([AB]?)/t$1/;
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "uplo" && do {
				$var = "ul";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "diag" && do {
				$var = "d";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "side" && do {
				$var = "s";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
		};
	}
	die "missed C parameters from '$func', '$paramList'" if scalar @processed != scalar @params;
	return join ", ", @processed;
}
