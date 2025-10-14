package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop/goods_gozero/goods"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()

	client := goods.NewGoodsClient(conn)

	// 1. 测试获取品牌列表
	fmt.Println("=== 测试获取品牌列表 ===")
	list, err := client.BrandList(context.Background(), &goods.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 10,
	})
	if err != nil {
		fmt.Printf("获取品牌列表失败: %v\n", err)
	} else {
		fmt.Printf("品牌列表: 总数=%d\n", list.Total)
		for i, brand := range list.Data {
			fmt.Printf("  %d. ID:%d 名称:%s Logo:%s\n", i+1, brand.Id, brand.Name, brand.Logo)
		}
	}
	fmt.Println("<UNK>")

	// 2. 测试创建品牌
	fmt.Println("\n=== 测试创建品牌 ===")
	createResp, err := client.CreateBrand(context.Background(), &goods.BrandRequest{
		Name: "测试品牌",
		Logo: "https://example.com/logo.png",
	})
	if err != nil {
		fmt.Printf("创建品牌失败: %v\n", err)
	} else {
		fmt.Printf("创建品牌成功: ID=%d 名称=%s Logo=%s\n",
			createResp.Id, createResp.Name, createResp.Logo)
	}

	// 3. 测试更新品牌（如果创建成功）
	if createResp != nil {
		fmt.Println("\n=== 测试更新品牌 ===")
		_, err = client.UpdateBrand(context.Background(), &goods.BrandRequest{
			Id:   createResp.Id,
			Name: "更新后的品牌",
			Logo: "https://example.com/new-logo.png",
		})
		if err != nil {
			fmt.Printf("更新品牌失败: %v\n", err)
		} else {
			fmt.Printf("更新品牌成功: ID=%d\n", createResp.Id)
		}

		// 4. 测试删除品牌
		fmt.Println("\n=== 测试删除品牌 ===")
		_, err = client.DeleteBrand(context.Background(), &goods.BrandRequest{
			Id: createResp.Id,
		})
		if err != nil {
			fmt.Printf("删除品牌失败: %v\n", err)
		} else {
			fmt.Printf("删除品牌成功: ID=%d\n", createResp.Id)
		}
	}

	// 5. 测试获取不存在的品牌
	fmt.Println("\n=== 测试获取不存在的品牌 ===")
	_, err = client.BrandList(context.Background(), &goods.BrandFilterRequest{
		Pages:       999,
		PagePerNums: 10,
	})
	if err != nil {
		fmt.Printf("获取不存在的品牌页面: %v\n", err)
	} else {
		fmt.Println("获取品牌列表成功")
	}
}
