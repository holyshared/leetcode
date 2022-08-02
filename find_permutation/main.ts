function findPermutation(s: string): number[] {
  let results: number[] = [];
  for (let i = 0; i < s.length + 1; i++) {
    results.push(0);
  }
  const stack: number[] = [];

  let j = 0;

  for (let i = 1; i <= s.length; i++) {
    if (s[i - 1] == "I") {
      stack.push(i);
      while (stack.length > 0) {
        results[j] = stack.pop()!;
        j++;
      }
    } else {
      stack.push(i);
    }
  }
  stack.push(s.length + 1);

  while (stack.length > 0) {
    results[j] = stack.pop()!;
    j++;
  }

  return results;
}
