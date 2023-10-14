function charToInt(c: string) {
    return c.charCodeAt(0) - 'a'.charCodeAt(0);
}

// Assume range is 26 lower case characters
// Theortical: Time - O(N); Memory - lower case of alphabetical
function isAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) return false;

    if (s.length === 1 && t.length === 1) {
        return s === t;
    }

    const a = new Uint8Array(26);
    const b = new Uint8Array(26);

    let i = 0;
    for ( ; i < s.length; i++) {
        a[charToInt(s[i])] += 1;
    }

    for (i = 0 ; i < t.length; i++) {
        b[charToInt(t[i])] += 1;
    }

    for (i = 0; i < a.length; i++) {
        if (a[i] !== b[i]) return false;
    }

    return true;
};


// Unicode implementation
function isAnagramU(s: string, t: string): boolean {
    const a = new Map<string, number>();
    const b = new Map<string, number>();

}

