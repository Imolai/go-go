#!/usr/bin/env perl
use feature qw(:5.10);
use strict;
use warnings;
use utf8;
use IO::File;
use Carp;
use English qw(-no_match_vars);
use Path::Tiny qw(path);

our $VERSION = 1.0.0;

my $input = 'README.src.md';
my $ifh   = IO::File->new($input, 'r');
defined $ifh || croak "$OS_ERROR: $input";
my $output = 'README.md';
my $ofh    = IO::File->new($output, 'w');
defined $ofh || croak "$OS_ERROR: $output";

say "Generating '$output' from '$input' and sources in 'src/'.";
foreach my $line ($ifh->getlines) {
  chomp $line;
  if ($line =~ m/^\[\[(.*)\]\]$/xms) {
    if (-e $1) {
      $ofh->print(path($1)->slurp);
    } else {
      croak $line;
    }
  } else {
    $ofh->print($line . "\n");
  }
}
undef $ifh;

system('doctoc', '--title', '## Contents', $output) == 0
  || croak "$CHILD_ERROR: doctoc";

