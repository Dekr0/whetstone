function topKFrequent(nums: number[], k: number): number[] {
    let b: number[][] = new Array(nums.length).fill([]);
    let m = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
        if (m.has(nums[i])) m.set(nums[i], m.get(nums[i]) as number + 1);
        else m.set(nums[i], 1);
    }

    m.forEach((v, k) => b[v].push(k));

    let r = [];
    for (let i = b.length - 1; i >= 0 && k > 0; i--) {
        if (b[i].length === 0) continue;
        for (let j = b[i].length; j >= 0 && k > 0; j--, k--) {
            r.push(b[i][j]);
        }
    }

    return r;
};

console.log(topKFrequent([1, 1, 1, 2, 2, 3], 2));
