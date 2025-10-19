var isSubsequence = function(s, t) {
    let sIndex = 0;
    let tIndex = 0;
    
    while (sIndex < s.length && tIndex < t.length) {    
        if (s[sIndex] === t[tIndex]) {
            sIndex++;
            tIndex++;
        } else {
            tIndex++;
        }
    }
    return sIndex === s.length;
}

var s = "abc", t = "ahbgdc";
console.log(isSubsequence(s, t)); // true
var s = "axc", t = "ahbgdc";
console.log(isSubsequence(s, t)); // false