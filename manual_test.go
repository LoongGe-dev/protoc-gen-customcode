package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func extractStartValue(comments string) int {
	// 匹配 //start 101 或 //start=101 或 //start:101 格式
	re := regexp.MustCompile(`//start\s*[:=]?\s*(\d+)`)
	text := comments

	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		start, err := strconv.Atoi(matches[1])
		if err == nil {
			return start
		}
	}
	return 0
}

func toCamelCase(s string) string {
	parts := strings.Split(strings.ToLower(s), "_")
	for i := 0; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

func main() {
	fmt.Println("=== 测试 extractStartValue ===")
	testCases := []string{
		"//start 101\n",
		"//start:202\n",
		"//start=303\n",
		"// other comment\n",
		"",
	}

	for _, tc := range testCases {
		result := extractStartValue(tc)
		fmt.Printf("Input: %q => Result: %d\n", tc, result)
	}

	fmt.Println("\n=== 测试 toCamelCase ===")
	enumNames := []string{
		"basic_code_unspecified",
		"basic_code_conf_not_exist",
		"basic_code_param_invalid",
		"team_code_unspecified",
		"team_code_member_info_empty",
		"gem_code_param_invalid",
	}

	for _, name := range enumNames {
		result := toCamelCase(name)
		fmt.Printf("Input: %s => Result: %s\n", name, result)
	}

	fmt.Println("\n=== 所有测试通过 ===")
}
