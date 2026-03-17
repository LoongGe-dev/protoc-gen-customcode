package main

import (
	"fmt"
	pb "github.com/LoongGe-dev/protoc-gen-customcode/examples/proto"
)

func main() {
	fmt.Println("=== 自定义错误码示例 ===\n")

	// 直接使用生成的常量
	fmt.Printf("CONF_NOT_EXIST 错误码: %d (101<<16 + 1)\n", pb.ConfNotExist_CustomCode)
	fmt.Printf("TEAM_USER_ID_INVALID 错误码: %d (102<<16 + 2)\n\n", pb.TeamUserIdInvalid_CustomCode)

	// 验证 GEM_PARAM_INVALID 的特殊需求
	gemParamInvalidCode := int32(102<<16 + 1)
	fmt.Printf("GEM_PARAM_INVALID 期望值: %d\n", gemParamInvalidCode)
	fmt.Printf("GEM_PARAM_INVALID 实际值: %d\n", pb.GemParamInvalid_CustomCode)
	fmt.Printf("匹配: %v\n\n", gemParamInvalidCode == pb.GemParamInvalid_CustomCode)

	// 使用错误码映射
	fmt.Println("=== 错误码映射 ===")
	testCodes := []int32{
		pb.ConfNotExist_CustomCode,
		pb.TeamUserIdInvalid_CustomCode,
		pb.GemParamInvalid_CustomCode,
		99999, // 不存在的错误码
	}

	for _, code := range testCodes {
		fmt.Printf("错误码 %d: %s\n", code, pb.GetErrorMessage(code))
	}

	fmt.Println("\n=== 错误码查询 ===")
	code := pb.GetErrorCode("TeamCode", "TEAM_FULL")
	fmt.Printf("TeamCode.TEAM_FULL 错误码: %d\n", code)

	fmt.Println("\n=== 验证计算逻辑 ===")
	fmt.Printf("BASIC_SUCCESS: %d = 101<<16 + 0\n", pb.BasicSuccess_CustomCode)
	fmt.Printf("TEAM_ALREADY_IN_TEAM: %d = 102<<16 + 3\n", pb.TeamAlreadyInTeam_CustomCode)
	fmt.Printf("GEM_SLOT_EMPTY: %d = 102<<16 + 4\n", pb.GemSlotEmpty_CustomCode)

	fmt.Println("\n=== 使用 ErrorCode 类型 ===")
	var errCode pb.ErrorCode = pb.ConfNotExist
	fmt.Printf("错误码类型: %T, 值: %d\n", errCode, errCode)
}
