function intervalIntersection(firstList: number[][], secondList: number[][]): number[][] {
  let i = 0, j = 0;
  let ans: number[][] = []

  while (i < firstList.length && j < secondList.length) {
    let lo = Math.max(firstList[i][0], secondList[j][0])
    let hi = Math.min(firstList[i][1], secondList[j][1])

    if (lo <= hi) {
      ans.push([lo, hi])
    }

    if (firstList[i][1] < secondList[j][1]) {
      i++
    } else {
      j++
    }
  }
  return ans
}
