function twoSum(nums: number[], target: number): number[] {
    const maps = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const r = target - nums[i];
        if (maps.has(r)) {
            const j = maps.get(r) as number;
            return [i, j];
        } else {
            maps.set(nums[i], i);
        }
    }

    return [];
};

console.log(twoSum([2,7,11,15], 9));
console.log(twoSum([3,2,4], 6));
console.log(twoSum([3,3], 6));
