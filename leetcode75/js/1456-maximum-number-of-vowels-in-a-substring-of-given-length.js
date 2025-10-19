var maxVowels = function(s, k) {
    const vowels = new Set(['a', 'e', 'i', 'o', 'u']);
    let maxCount = 0;
    let currentCount = 0;

    // Count vowels in the first window of size k
    for (let i = 0; i < k; i++) {
        if (vowels.has(s[i])) {
            currentCount++;
        }
    }
    maxCount = currentCount;

    // Slide the window over the string
    for (let i = k; i < s.length; i++) {
        if (vowels.has(s[i])) {
            currentCount++;
        }
        if (vowels.has(s[i - k])) {
            currentCount--;
        }
        maxCount = Math.max(maxCount, currentCount);
    }

    return maxCount;
};

var s = "abciiidef", k = 3
console.log(maxVowels(s, k)) // 3
var s = "aeiou", k = 2
console.log(maxVowels(s, k)) // 2
var s = "leetcode", k = 3
console.log(maxVowels(s, k)) // 2