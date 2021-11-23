package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	num := []byte(num1)
	ans := "0"
	for i:=len(num2)-1;i>=0;i--{
		if i < len(num2)-1 {
			num = append(num, '0')
		}
		ans = plus(ans, mult(string(num), num2[i]))
	}
	return ans
}

func plus(bs1 string, bs2 string) string {
	num1,num2 := reverse([]byte(bs1)),reverse([]byte(bs2))
	ans := make([]byte, 0, len(bs1))
	n1,n2:=len(num1),len(num2)
	v := 0
	for i:=0;i<n1&&i<n2;i++{
		v1 := num1[i] - '0'
		v2 := num2[i] - '0'
		v = int(v1 + v2) + v
		ans = append(ans, byte(v % 10) + '0')
		v = v / 10
	}
	if n1 > n2 {
		for i:=n2;i<n1;i++{
			v += int(num1[i] - '0')
			ans = append(ans, byte(v % 10) + '0')
			v = v / 10
		}
	}

	if n2 > n1 {
		for i:=n1;i<n2;i++{
			v += int(num2[i] - '0')
			ans = append(ans, byte(v % 10) + '0')
			v = v / 10
		}
	}

	for v > 0 {
		ans = append(ans, byte(v % 10) + '0')
		v /= 10
	}
	return string(reverse(ans))
}

func mult(bs1 string, num byte) string {
	num1 := reverse([]byte(bs1))
	vv := int(num - '0')
	v := 0
	ans := make([]byte, 0, len(bs1)+2)
	for i:=0;i<len(num1);i++{
		v += int(num1[i]-'0') * vv
		ans = append(ans, byte(v % 10) + '0')
		v = v / 10
	}
	for v > 0 {
		ans = append(ans, byte(v % 10) + '0')
		v = v / 10
	}
	return string(reverse(ans))
}

func reverse(b []byte) []byte {
	for i,j:=0,len(b)-1;i<j;i,j=i+1,j-1 {
		b[i],b[j] = b[j],b[i]
	}
	return b
}





func combinationSum2(candidates []int, target int) [][]int {
	freq := map[int]int{}
	for _, v := range candidates {
		freq[v]++
	}

	vv := make([]int, 0, len(freq))
	for v := range freq {
		vv = append(vv, v)
	}


	ans := [][]int{}
	var dfs func(target int, idx int, ints []int)
	dfs = func(target int, idx int, ints []int) {
		if target == 0 {
			ans = append(ans, append([]int{}, ints...))
			return
		}
		if target < 0 || idx >= len(vv) {
			return
		}

		// 不选
		dfs(target, idx+1, ints)
		// 选择
		num := freq[vv[idx]]
		for i:=1;i<=num;i++{
			if target - i * vv[idx] >= 0 {
				ints = append(ints, vv[idx])
				dfs(target - i * vv[idx], idx+1, ints)
			}
		}
	}
	dfs(target, 0, nil)
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	var dfs func(target int, idx int, v []int)
	dfs = func(target int, idx int, v []int) {
		if target == 0 {
			vv := make([]int, len(v))
			copy(vv, v)
			ans = append(ans, vv)
			return
		}
		if target < 0 || idx >= len(candidates) {
			return
		}

		dfs(target, idx+1, v)
		v = append(v, candidates[idx])
		dfs(target - candidates[idx], idx, v)
	}
	dfs(target, 0, nil)
	return ans
}

func countAndSay(n int) string {
	ans := make([]string, n+1)
	ans[1] = "1"
	for i:=2;i<=n;i++{
		cnt := 1
		c := ans[i-1][0]
		for j := 1; j < len(ans[i-1]);j++{
			if ans[i-1][j] != c {
				ans[i] += strconv.Itoa(cnt) + string([]byte{c})
				c = ans[i-1][j]
				cnt = 1
			} else {
				cnt++
			}
		}
		if cnt > 0 {
			ans[i] += strconv.Itoa(cnt) + string([]byte{c})
		}
	}
	return ans[n]
}

func isMatch(s string, p string) bool {
	m,n:=len(s),len(p)
	dp := make([][]bool, m+1)
	for i:=0;i<=m;i++{
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i:=0;i<=m;i++{
		for j:=1;j<=n;j++{
			if p[j-1] == '*' {
				if i > 0 && match(s[i-1],p[j-2]) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else if i > 0 && match(s[i-1],p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func match(s byte, d byte) bool {
	return d == '.' || s == d
}


func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for first := 0; first < len(nums)-3;first++{
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first+1;second<len(nums)-2;second++{
			if second > first+1 && nums[first] == nums[second] {
				continue
			}
			third := second + 1
			four := len(nums)-1
			v := target - nums[first] - nums[second]
			for third < four {
				if nums[third] + nums[four] == v {
					ans = append(ans, []int{nums[first],nums[second],nums[third], nums[four]})
					for third < four && nums[four] == nums[four-1] {
						four--
					}
					for third < four && nums[third] == nums[third+1] {
						third++
					}
					four--
					third++
				} else if nums[third] + nums[four] > v {
					four--
				} else {
					third++
				}
			}
		}
	}
	return ans
}



func myAtoi(s string) int {
	bits := []byte(s)
	startIndex := 0
	for startIndex<len(bits) && bits[startIndex] == ' '{
		startIndex++
	}
	flag := 0
	if startIndex < len(bits) {
		if bits[startIndex] == '-' {
			flag = 1
			startIndex++
		} else if bits[startIndex] == '+' {
			startIndex++
		}
	}
	ans := 0
	if flag == 1 && startIndex < len(bits) {
		if bits[startIndex] >= '0' && bits[startIndex] <= '9' {
			ans = -int(bits[startIndex] - '0')
			startIndex++
		}
	}
	for i:=startIndex;i<len(bits);i++{
		if bits[i] < '0' || bits[i] > '9' {
			break
		}
		if ans > math.MaxInt32 {
			return math.MaxInt32
		}
		if ans < math.MinInt32 {
			return math.MinInt32
		}
		ans = ans * 10 + int(bits[i] - '0')
	}
	return ans
}



var(
	m sync.Map
)
func Caller(skip int)(pc uintptr, file string, line int, ok bool){
	rpc := [1]uintptr{}
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	var (
		frame  runtime.Frame
	)
	pc  = rpc[0]
	if item,ok:=m.Load(pc);ok{
		frame = item.(runtime.Frame)
	}else{
		tmprpc := []uintptr{
			pc,
		}
		frame, _ = runtime.CallersFrames(tmprpc).Next()
		m.Store(pc,frame)
	}
	return frame.PC,frame.File,frame.Line,frame.PC!=0
}


func minTimeToType(word string) int {
	ans := 0
	c := byte('a')
	v := 0
	for i:=0;i<len(word);i++{
		if word[i] > c {
			v = int('z' - word[i]) + int(c - 'a') + 1
			x := int(word[i] - c)
			if v > x {
				v = x
			}
		} else {
			v = int('z' - c) + int(word[i] - 'a') + 1
			x := int(c - word[i])
			if v > x {
				v = x
			}
		}
		c = word[i]
		ans += v+1
	}
	return ans
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i:=0;i<n;i++{
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for i:=0;i<n-1;i++ {
		dp[i][i+1] = s[i] == s[i+1]
	}
	for i:=n-3;i>=0;i--{
		for j:=i+2;j<n;j++{
			dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
		}
	}
	maxLen := 0
	ans := ""
	for i:=0;i<n;i++{
		for j:=i;j<n;j++{
			if dp[i][j] && j-i+1>maxLen {
				maxLen = j-i+1
				ans = s[i:j+1]
			}
		}
	}
	return ans
}



func checkRecord(n int) int {
	// dp[i][j][k] 第i天 缺勤i次数 后缀连续k
	dp := make([][][]int, n+1)
	for i:=0;i<n;i++{
		dp[i] = make([][]int, 2)
		for j:=0;j<2;j++{
			dp[i][j] = make([]int, 3)
		}
	}
	mod := int(1e9+7)
	dp[0][0][0] = 1
	for i:=1;i<=n;i++{
		// 最后一个为P
		for j:=0;j<2;j++{
			for k:=0;k<3;k++{
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}
		// 最后一个为A
		for k:=0;k<3;k++{
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}
		// 最后一个为L
		for j:=0;j<2;j++{
			for k:=0;k<2;k++{
				dp[i][j][k+1] = (dp[i][j][k+1] + dp[i-1][j][k]) % mod
			}
		}
	}

	ans := 0
	for j:=0;j<2;j++{
		for k:=0;k<3;k++ {
			ans = (ans + dp[n][j][k]) % mod
		}
	}
	return ans
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	bits := strconv.Itoa(n)

	dp := make([]int, 11)
	v := 1
	for i:=1;i<11;i++{
		dp[i] = dp[i-1]*10 + v
		v *= 10
	}

	n = len(bits)
	first,sencond := 0, 0
	if int(bits[len(bits)-1]-'0') >= 1 {
		first = 1
	}
	mult := 1
	v = int(bits[len(bits)-1] - '0')
	for i:=n-2;i>=0;i--{
		mult *= 10
		v = int(bits[i] - '0') * mult + v
		if bits[i] == '0' {
			sencond = first
		} else if bits[i] == '1' {
			sencond = dp[n-1-i] + first + (v % mult) + 1
		} else {
			sencond = dp[n-1-i] * int(bits[i]-'0') + first + mult
		}
		first = sencond
	}
	return first
}

func nthUglyNumber(n int, a int, b int, c int) int {
	dp := make([]int, n+1)
	primes := []int{a,b,c}
	points := []int{0,0,0}
	dp[0] = 1
	for i:=1;i<=n;i++{
		dp[i] = dp[points[0]] * primes[0]
		for j:=1;j<len(primes);j++{
			v := dp[points[j]] * primes[j]
			if dp[i] > v {
				dp[i] = v
			}
		}

		for j:=0;j<len(primes);j++{
			v := dp[points[j]] * primes[j]
			if dp[i] == v {
				points[j]++
			}
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	uparr := make([]int, 0, len(obstacles))
	for i:=0;i<len(obstacles);i++{
		pos := sort.Search(len(uparr), func(j int) bool {
			return uparr[j] > obstacles[i]
		})
		if pos == len(uparr) {
			uparr = append(uparr, obstacles[i])
		} else {
			uparr[pos] = obstacles[i]
		}
		ans[i] = pos + 1
	}
	return ans
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	var dfs func(clone []int, idx int) int

	check := func(needs []int, special []int) bool {
		for i:=0;i<len(needs);i++ {
			if needs[i] < special[i] {
				return false
			}
		}
		return true
	}

	dfs = func(needs []int, idx int) int {
		clone := make([]int, len(needs))
		copy(clone, needs)

		ans := 0
		for i:=0;i<len(clone);i++{
			ans += clone[i] * price[i]
		}


		if idx >= len(special) {
			return ans
		}

			// 不购买
			ans = min(ans, dfs(clone, idx+1))
			if check(clone, special[idx]) {
				// 购买
				for j:=0;j<len(clone);j++{
					clone[j] -=  special[idx][j]
				}
				ans = min(ans, special[idx][len(special[idx])-1] + dfs(clone, idx))
			}

		return ans
	}
	return dfs(needs, 0)
}

func min(left,right int) int {
	if left < right {
		return left
	}
	return right
}

func countArrangement(n int) int {
	mask := (1<<n) - 1
	ans := 0
	var dfs func(state int, k int)
	dfs = func(state int, k int) {
		if k == n {
			if state == mask {
				ans += 1
			}
			return
		}
		for i:=0;i<n;i++{
			if (state >> i) & 1 == 1 {
				continue
			}
			if (i+1) % k == 0 || k % (i+1) == 0 {
				dfs(state | (1<<i), k+1)
			}
		}
	}
	dfs(0, 1)
	return ans
}

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(i int) int {
		return (i  + nums[i] % n + n) % n
	}
	for i:=0;i<n;i++{
		slow,fast:=i,next(i)
		for slow != fast {
			if nums[slow] * nums[fast] < 0 || nums[slow] * nums[next(fast)] < 0 || nums[slow] * nums[i] < 0 {
				break
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		if slow == fast && next(slow) != i {
			return true
		}
	}
	return false
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	dist := make([][]int, n)
	for i:=0;i<n;i++{
		dist[i] = make([]int, n)
		for j:=0;j<n;j++{
			dist[i][j] = (1<<20)
		}
	}
	for i, g := range graph {
		for _, j := range g {
			dist[i][j] = 1
			dist[j][i] = 1
		}
	}
	for k:=0;k<n;k++{
		for i:=0;i<n;i++{
			for j:=0;j<n;j++{
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	mask := (1<<n)
	dp := make([][]int, n)
	for i:=0;i<n;i++{
		dp[i] = make([]int, mask)
		for j:=0;j<mask;j++{
			dp[i][j] = (1<<20)
		}
		dp[i][(1<<i)] = 0
	}

	for state := 1; state < mask; state++ {
		for i:=0;i<n;i++{
			if (state >> i) & 1 == 0 {
				continue
			}
			for j:=0;j<n;j++{
				if (state >> j) & 1 == 1 {
					continue
				}
				dp[j][state|(1<<j)] = min(dp[j][state|(1<<j)], dp[i][state]+dist[i][j])
			}
		}
	}
	ans := (1<<20)
	for i:=0;i<n;i++{
		ans = min(ans, dp[i][mask-1])
	}
	return ans
}

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i,j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	idx := sort.Search(len(ranges), func(i int) bool {
		return ranges[i][0] >= left
	})

	if idx > 0 {
		idx--
	}

	for k := left; k <= right; k++ {
		for ;idx<len(ranges);idx++{
			if ranges[idx][0] <= k && k <= ranges[idx][1] {
				break
			}
			if ranges[idx][0] > k || idx == len(ranges)-1 {
				return false
			}
		}
	}
	return true
}


type LRUCache struct {
	ll list.List
	num int
	capacity int
	vals map[int]*list.Element
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:capacity,
		vals:map[int]*list.Element{},
	}
}


func (this *LRUCache) Get(key int) int {
	if v, ok := this.vals[key]; ok {
		this.ll.MoveToBack(v)
		return v.Value.([]int)[1]
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.vals[key]; ok {
		v.Value.([]int)[1] = value
		this.ll.Remove(v)
		this.ll.PushBack(v)
		return
	}

	if this.num >= this.capacity {
		v := this.ll.Remove(this.ll.Front())
		delete(this.vals, v.([]int)[0])
		this.num--
	}

	this.ll.PushBack([]int{key, value})
	this.vals[key] = this.ll.Back()
	this.num++
}

func waysToSplit(nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i:=0;i<len(nums);i++{
		preSum[i+1] = preSum[i] + nums[i]
	}
	var ans int
	for k := 0; k < len(nums)-2; k++{
		left := preSum[k+1]
		preSum1 := preSum[k+1:]
		idxLeft := sort.Search(len(preSum1), func(i int) bool {
			return preSum1[i+1]  >= (left << 1)
		})
		idxRight := sort.Search(len(preSum1), func(i int) bool {
			return preSum1[len(preSum1)-1] - preSum1[i+1] < preSum1[i+1] - left
		})

		if idxRight - idxLeft > 0 {
			ans += idxRight - idxLeft
		} else {
			break
		}
	}
	return ans % (1e9+7)
}

func maxRepOpt1(text string) int {
	ans := 0
	left := 0
	changeNum := 0
	exists := make([]bool, 26)
	c := text[left]
	for i:=1;i<len(text);i++{
		if text[i] == c {
			if i == len(text) - 1 {
				ans = max(ans, i - left + 1)
			}
			continue
		}
		if changeNum == 0 && exists[c-'a'] {
			if i == len(text) - 1 {
				ans = max(ans, i - left + 1)
			}
			changeNum++
			continue
		}

		ans = max(ans, i-left)
		for (changeNum > 0 || !exists[c-'a']) && left < i {
			if text[left] != c {
				changeNum--
			}
			exists[text[left]-'a'] = true
			left++
		}
		c = text[left]
		if c != text[i] {
			i--
		}
	}

	right := len(text)-1
	changeNum = 0
	exists = make([]bool, 26)
	c = text[right]
	for i := len(text)-2; i >= 0; i--{
		if text[i] == c {
			if i == 0 {
				ans = max(ans, right - i + 1)
			}
			continue
		}
		if changeNum == 0 && exists[c-'a'] {
			if i == 0 {
				ans = max(ans, right - i + 1)
			}
			changeNum++
			continue
		}

		ans = max(ans, right - i)
		for (changeNum > 0 || !exists[c-'a']) && right > i {
			if text[right] != c {
				changeNum--
			}
			exists[text[right]-'a'] = true
			right--
		}
		c = text[right]
	}
	return ans
}

func max(left,right int) int {
	if left < right {
		return right
	}
	return left
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	nums := append([]int{}, nums1...)
	sort.Ints(nums)
	sum := 0
	maxx := 0
	for i:=0;i<len(nums1);i++{
		diff := abs(nums1[i]-nums2[i])
		sum += diff
		idx := sort.Search(len(nums), func(k int) bool {
			return nums[k] >= nums2[i]
		})
		if idx < len(nums) {
			v := nums[idx] - nums2[i]
			if v < diff {
				maxx = max(diff - v, maxx)
			}
		}
		if idx > 0 {
			v := nums2[i]-nums[idx-1]
			if v < diff {
				maxx = max(diff - v, maxx)
			}
		}
	}
	return (sum - maxx) % (1e9+7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}




func findAnagrams(s string, p string) []int {
	valid := 0
	ori := map[byte]int{}
	target := map[byte]int{}
	left := 0
	ans := make([]int, 0, 8)
	for i:=0;i<len(p);i++ {
		ori[p[i]]++
	}

	for i:=0;i<len(s);i++{
		if _, ok := ori[s[i]]; ok {
			target[s[i]]++
			if target[s[i]] == ori[s[i]] {
				valid++
			}
		} else {
			left = i + 1
			target = map[byte]int{}
			valid = 0
			continue
		}
		if i - left + 1 == len(p) {
			if valid == len(ori) {
				ans = append(ans, left)
			}
			if target[s[left]] == ori[s[left]] {
				valid--
			}
			target[s[left]]--
			left++
		}
	}
	return ans
}

func lengthOfLongestSubstring(s string) int {
	ans, left := 0, 0
	vals := make([]uint64, 2)
	for i:=0;i<len(s);i++{
		v := uint64(s[i])
		idx := v / 64
		mask := uint64(1) << (v-idx*64)
		for vals[idx] & mask > 0 {
			vv := s[left]
			j := vv / 64
			vals[j] &= ^(1 << (vv-j*64))
			left++
		}
		vals[idx] |= mask
		if ans <= i - left + 1 {
			ans = i - left + 1
		}
	}
	return ans
}

func longestNiceSubstring(s string) string {
	ans := []int{-1,-1}
	for k := 1; k <= 26; k++ {
		low := make([]int, 26)
		up := make([]int, 26)
		left := 0
		ktypes := 0
		for i:=0;i<len(s); i++ {
			if s[i]>='A' && s[i] <= 'Z' {
				if up[s[i]-'A'] + low[s[i]-'A'] == 0 {
					ktypes++
				}
				up[s[i]-'A']++
			} else {
				if up[s[i]-'a'] + low[s[i]-'a'] == 0 {
					ktypes++
				}
				low[s[i]-'a']++
			}
			for ktypes > k {
				if s[left] >= 'A' && s[left] <= 'Z' {
					up[s[left]-'A']--
					if up[s[left]-'A'] + low[s[left]-'A'] == 0 {
						ktypes--
					}
				} else {
					low[s[left]-'a']--
					if up[s[left]-'a'] + low[s[left]-'a'] == 0 {
						ktypes--
					}
				}
				left++
			}
			check := true
			for i:=0;i<len(low);i++{
				if (low[i] > 0 && up[i] <= 0) || (low[i] <= 0 && up[i] > 0) {
					check = false
					break
				}
			}
			if check {
				if ans[0] == -1 || ans[1] - ans[0] < i - left {
					ans[0],ans[1] = left, i
				}
			}
		}
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[0]:ans[1]+1]
}



func findMinHeightTrees(n int, edges [][]int) []int {
	es := make([][]int, n + 1)
	degree := make([]int, n+1)
	for _, e := range edges {
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}
	queue := make([]int, 0, 32)
	for i:=1;i<=n;i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		size := len(queue)
		if size == n {
			break
		}
		n -= size
		for i:=0;i<size; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, e := range es[curr] {
				if degree[e] == 1 {
					continue
				}
				degree[e]--
				if degree[e] == 1 {
					queue = append(queue, e)
				}
			}
		}
	}
	return queue
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	cnts := make([]int, 1e5+1)
	for i:=0;i<len(nums);i++{
		cnts[nums[i]]++
	}
	ans := 0
	for i:=len(nums)-1;i>=0;i--{
		i = i - cnts[nums[i]] + 1
		v := k
		cnt := cnts[nums[i]]
		for j := i - 1; j>=0;j--{
			j = j - cnts[nums[j]] + 1
			diff := nums[i]-nums[j]
			if cnts[nums[j]] * diff <= v {
				cnt += cnts[nums[j]]
				v -= cnts[nums[j]] * diff
			} else {
				cnt += v / diff
				break
			}
		}
		ans = max(ans, cnt)
	}
	return ans
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	sort.Slice(rooms, func(i,j int)bool{
		return rooms[i][1] < rooms[j][1]
	})

	ans := make([]int, len(queries))
	for k, q := range queries {
		idx := sort.Search(len(rooms), func(i int) bool {
			return rooms[i][1] >= q[1]
		})
		ans[k] = -1
		for i:=idx;i<len(rooms);i++{
			if ans[k] < 0 || abs(rooms[i][0]-q[0]) < ans[k] {
				ans[k] = rooms[i][0]
			}
		}
	}
	return ans
}


func permutation(s string) []string {
	nums := []byte(s)
	m := map[string]struct{}{}
	m[string(nums)] = struct{}{}
	for {
		nextPermutation(nums)
		if _, ok := m[string(nums)]; ok {
			break
		}
	}
	ans := make([]string, 0, len(m))
	for str := range m {
		ans = append(ans, str)
	}
	return ans
}

func nextPermutation(nums []byte) {
	i,j,k:=len(nums)-2,len(nums)-1,len(nums)-1
	for i>=0 && nums[i] >= nums[j] {
		i,j = i-1,j-1
	}
	if i >= 0 {
		for nums[k] <= nums[i] {
			k--
		}
		nums[k], nums[i] = nums[i], nums[k]
	}
	for i,j:=j,len(nums)-1; i<j; i,j = i+1,j-1{
		nums[i],nums[j] = nums[j],nums[i]
	}
}



func numsGame(nums []int) []int {
	low := MinIntHeap{}
	high := MaxIntHeap{}
	for i:=0;i<len(nums);i++{
		nums[i] -= i
	}
	lowSum := 0
	highSum := 0
	ans := make([]int, len(nums))
	for i:=0;i<len(nums);i++{
		heap.Push(&low, nums[i])
		lowSum += nums[i]
		top := low[0]
		lowSum -= top
		heap.Push(&high, heap.Pop(&low))
		highSum += top
		if len(low) < len(high) {
			top := high[0]
			highSum -= top
			lowSum += top
			heap.Push(&low, heap.Pop(&high))
		}

		if (i+1)&1==1{
			mid := low[0]
			lowSize,highSize:=len(low),len(high)
			ans[i] = (mid * lowSize - lowSum + highSum - highSize * mid) % (1e9+7)
		} else {
			mid := low[0]
			lowSize,highSize:=len(low),len(high)
			ans1 := mid * lowSize - lowSum + highSum - highSize * mid

			mid = high[0]
			ans2 := mid * lowSize - lowSum + highSum - highSize * mid
			ans[i] = ans1 % (1e9+7)
			if ans1 > ans2 {
				ans[i] = ans2 % (1e9+7)
			}
		}
	}
	return ans
}

type MinIntHeap []int
func (a *MinIntHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}
func (a *MinIntHeap) Pop() interface{} {
	v := (*a)[len(*a)-1]
	*a=(*a)[:len(*a)-1]
	return v
}
func (a *MinIntHeap) Len() int {
	return len(*a)
}
func (a *MinIntHeap) Less(i, j int) bool {
	return (*a)[i] < (*a)[j]
}
func (a *MinIntHeap) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

type MaxIntHeap []int
func (a *MaxIntHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}
func (a *MaxIntHeap) Pop() interface{} {
	v := (*a)[len(*a)-1]
	*a=(*a)[:len(*a)-1]
	return v
}
func (a *MaxIntHeap) Len() int {
	return len(*a)
}
func (a *MaxIntHeap) Less(i, j int) bool {
	return (*a)[i] > (*a)[j]
}
func (a *MaxIntHeap) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func numberOfRounds(startTime string, finishTime string) int {
	strs := strings.Split(startTime, ":")
	v1, _ := strconv.Atoi(strs[0])
	v2, _ := strconv.Atoi(strs[1])
	stime := v1*60 + v2
	strs = strings.Split(finishTime, ":")
	v1, _ = strconv.Atoi(strs[0])
	v2, _ = strconv.Atoi(strs[1])
	etime := v1 * 60 + v2
	if etime <= stime {
		etime += 24*60
	}
	if stime % 15 != 0 {
		stime += 15 - (stime % 15)
	}
	if etime % 15 != 0 {
		etime -= etime % 15
	}
	return (etime - stime)/15
}

func largestArea(g []string) int {
	m := len(g)
	n := len(g[0])

	grid := make([][]byte, m)
	for i:=0;i<m;i++{
		grid[i] = []byte(g[i])
	}

	ans := 0
	stack := make([]Pair, 0, 16)
	for i := 0;i < m;i++{
		for j:=0;j<n;j++{
			if grid[i][j] == '0' || grid[i][j] == '*' {
				continue
			}
			c := grid[i][j]
			valid := true
			grid[i][j] = '*'
			stack = append(stack, Pair{First:i,Second:j})
			val := 1
			for len(stack) > 0 {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if v.First == 0 || v.First == m-1 || v.Second == 0 || v.Second == n - 1 {
					valid = false
				}
				ii,jj := v.First - 1,v.Second
				if ii >= 0 && ii < m && jj >= 0 && jj < n {
					if grid[ii][jj] == c {
						grid[ii][jj] = '*'
						val++
						stack = append(stack, Pair{First:ii,Second:jj})
					}
					if grid[ii][jj] == '0' {
						valid = false
					}
				}

				ii,jj = v.First + 1,v.Second
				if ii >= 0 && ii < m && jj >= 0 && jj < n {
					if grid[ii][jj] == c {
						grid[ii][jj] = '*'
						val++
						stack = append(stack, Pair{First:ii,Second:jj})
					}
					if grid[ii][jj] == '0' {
						valid = false
					}
				}

				ii,jj = v.First ,v.Second-1
				if ii >= 0 && ii < m && jj >= 0 && jj < n {
					if grid[ii][jj] == c {
						grid[ii][jj] = '*'
						val++
						stack = append(stack, Pair{First:ii,Second:jj})
					}
					if grid[ii][jj] == '0' {
						valid = false
					}
				}

				ii,jj = v.First,v.Second+1
				if ii >= 0 && ii < m && jj >= 0 && jj < n {
					if grid[ii][jj] == c {
						grid[ii][jj] = '*'
						val++
						stack = append(stack, Pair{First:ii,Second:jj})
					}
					if grid[ii][jj] == '0' {
						valid = false
					}
				}
			}
			if valid {
				ans = max(ans, val)
			}
		}
	}
	return ans
}


type Pair struct {
	First int
	Second int
}

func maxLength(arr []string) int {
	masks := make([]int, 0, len(arr))
	for i:=0;i<len(arr);i++ {
		mask := 0
		for j:=0;j<len(arr[i]);j++{
			idx := int(arr[i][j] - 'a')
			if mask & (1<<idx) == 1 {
				mask = -1
				break
			}
			mask |= (1<<idx)
		}
		if mask > 0 {
			masks = append(masks, mask)
		}
	}

	n := len(masks)
	var ans int
	var dfs func(pos int, mask int)
	dfs = func(pos int, mask int) {
		if pos == n {
			cnt := 0
			for mask > 0 {
				cnt++
				mask &= mask - 1
			}
			ans = max(ans, cnt)
			return
		}
		// 选pos
		if mask ^ masks[pos] == mask | masks[pos] {
			dfs(pos+1, mask | masks[pos])
		}
		// 不选pos
		dfs(pos+1,mask)
	}
	dfs(0, 0)
	return ans
}

func canReach(s string, minJump int, maxJump int) bool {
	dp := make([]int, len(s)+1)
	sum := make([]int, len(s)+2)
	dp[1] = 1
	for i:=1;i<=minJump;i++ {
		sum[i] = sum[i-1]+dp[i-1]
	}
	for i:=minJump+1;i<=len(s);i++{
		if s[i-1] == '0' {
			l,r := i-maxJump,i-minJump
			if l <= 1 {
				l = 1
			}
			dp[i] = sum[r] - sum[l-1]
		}
		sum[i+1] = sum[i]+dp[i]
	}
	return dp[len(s)] > 0
}

func minSpeedOnTime(dist []int, hour float64) int {
	eps := 1e-7
	sum := 0
	for i:=0;i<len(dist);i++{
		sum += dist[i]
	}
	maxSpeed := int(float64(sum) / hour) + 1
	for i:=1; i <= maxSpeed; i++ {
		cost := float64(0)
		for j:=0;j<len(dist)-1;j++ {
			cost += math.Ceil(float64(dist[j])/float64(i) - eps)
		}
		cost += float64(dist[len(dist)-1])/float64(i)
		if cost <= hour {
			return i
		}
	}
	return -1
}


func minimumXORSum(nums1 []int, nums2 []int) int {
	ans := math.MaxInt64
	dfs(nums1, 0, nums2, 0, 0, &ans)
	return ans
}

func dfs(nums1 []int, idx int, nums2 []int, idxMap int, ans int, minAns *int) {
	if idx == len(nums1) {
		if *minAns > ans {
			*minAns = ans
		}
		fmt.Println("ans=", ans)
		return
	}
	if ans >= *minAns {
		return
	}
	for i:=0;i<len(nums2);i++{
		if idxMap & (1<<i) == 0 {
			tmp := ans + nums1[idx]^nums2[i]
			fmt.Println(idx, i, ans, nums1[idx], nums2[i],nums1[idx]^nums2[i], tmp)
			dfs(nums1, idx+1, nums2, idxMap|(1<<i), ans + nums1[idx]^nums2[i], minAns)
		}
	}
}