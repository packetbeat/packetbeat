#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# Generate system call table for OpenBSD from master list
# (for example, /usr/src/sys/kern/syscalls.master).

use strict;

if($ENV{'GOARCH'} eq "" || $ENV{'GOOS'} eq "") {
	print STDERR "GOARCH or GOOS not defined in environment\n";
	exit 1;
}

my $command = "mksysnum_netbsd.pl " . join(' ', @ARGV);

print <<EOF;
// $command
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build $ENV{'GOARCH'},$ENV{'GOOS'}

package unix

const (
EOF

my $line = '';
while(<>){
	if($line =~ /^(.*)\\$/) {
		# Handle continuation
		$line = $1;
		$_ =~ s/^\s+//;
		$line .= $_;
	} else {
		# New line
		$line = $_;
	}
	next if $line =~ /\\$/;
	if($line =~ /^([0-9]+)\s+((STD)|(NOERR))\s+(RUMP\s+)?({\s+\S+\s*\*?\s*\|(\S+)\|(\S*)\|(\w+).*\s+})(\s+(\S+))?$/) {
		my $num = $1;
		my $proto = $6;
		my $compat = $8;
		my $name = "$7_$9";

		$name = "$7_$11" if $11 ne '';
		$name =~ y/a-z/A-Z/;

		if($compat eq '' || $compat eq '13' || $compat eq '30' || $compat eq '50') {
			print "	$name = $num;  // $proto\n";
		}
	}
}

print <<EOF;
)
EOF
