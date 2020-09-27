/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func search(start, end int) int {
		half := int((end - start)/2) + start
		if !isBadVersion(half) {
				return search(half+1, end)
		}

		if isBadVersion(half-1) {
				return search(start, half)
		}

		return half

}

func firstBadVersion(n int) int {
		return search(1, n+1)
}
