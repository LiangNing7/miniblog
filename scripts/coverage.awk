#!/usr/bin/env awk

{
  print $0Add commentMore actions
  if (match($0, /^total:/)) {
    sub(/%/, "", $NF);
    printf("test coverage is %s% (quality gate is %s%)\n", $NF, target)
    if (strtonum($NF) < target) {
      printf("test coverage does not meet expectations: %d%, please add test cases!\n", target)
      exit 1;
    }
  }
}
