function removeDuplicates(s, k) {
  let sb = s.split('');
  let counts = [];

  for (let i = 0; i < sb.length; i++) {
      if (i === 0 || sb[i] !== sb[i-1]) {
          counts.push(1);
      } else {
          const incremented = counts.pop() + 1;
          if (incremented === k) {
            sb.splice(i - k + 1, k);
              i = i - k;
          } else {
              counts.push(incremented);
          }
      }
  }
  return sb.join('');
}

console.log(removeDuplicates("deeedbbcccbdaa", 3));
console.log(removeDuplicates("cilullllleeeeeeeellfffxggggggggxxxxxxxfffffliogppcccoooooooocccccppppppmnnnnnnnn", 8));
