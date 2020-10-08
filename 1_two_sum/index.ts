function twoSum(nums: number[], target: number): number[] {
  const values = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    const index = values.get(diff);
    if (index !== undefined && index !== i) {
      return [index, i];
    }
    values.set(nums[i], i);
  }
  return [];
}

console.log(twoSum([2, 7, 11, 15], 9));
console.log(twoSum([3, 2, 4], 6));
console.log(twoSum([3, 3], 6));
console.log(twoSum([3, 2, 1], 3));
