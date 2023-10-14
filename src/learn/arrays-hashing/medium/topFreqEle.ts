function topKFrequent(nums: number[], k: number): number[] {
    const buckets: Array<Set<number>> = new Array<Set<number>>(nums.length + 1); // 0 ~ nusm.length
    const map: Map<number, number> = new Map<number, number>();

    let i = 0;
    // Initialize buckets
    for (; i < buckets.length; i++) {
        buckets[i] = new Set<number>();
    }

    // Count all appearance
    nums.forEach(num => {
        const count = map.get(num);
        if (!count) {
            map.set(num, 1);
        } else {
            map.set(num, count + 1);
        }
    });
    
    // Write into buckets
    map.forEach((value, key) => {
        buckets[value].add(key);
    });

    let result: number[] = [];
    // Collect reuslt, k is guaranteed that 1 <= k <= nums.length
    for (i = buckets.length - 1; i > 0 && k > 0; i--) {
        if (buckets[i].size === 0) continue;

        buckets[i].forEach(num => { result.push(num); k--; });
    }

    return result;
}

function topKFrequentFast(nums: number[], k: number): number[] {
    // use a map instead of object, since the key needs to be a number.
    // Objects can only use strings or symbols as keys;
    const countMap = new Map<number,number>();
    nums.forEach((num) => {
        countMap.set(num, (countMap.get(num) || 0) + 1);
    });
 
    const buckets: number[][] = [];
    // add counts into buckets
    countMap.forEach((count, num) => {
        buckets[count] = (buckets[count] || []);
        buckets[count].push(num)
    })

    // loop thru buckets in reverse, to get most frequent.
    // build up a return array until we hit length of k
    let results = [];
    for (let i = buckets.length - 1; i > 0; i--) {
        const nums = buckets[i];
        if (nums) {
            results.push(...nums)
        }
        if (results.length === k) {
            return results;
        }
    }

    return results;
};
