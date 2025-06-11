// ElasticView插件HelloWorld API接口文件 - 定义与后端HelloWorld功能相关的HTTP请求
// 使用ElasticView插件SDK的request方法与插件后端进行通信

import {request} from '@elasticview/plugin-sdk' // 导入ElasticView插件SDK的HTTP请求函数

/**
 * HelloAction API接口 - 测试插件与ElasticSearch连接的连通性
 * 
 * 功能说明：
 * - 调用插件后端的HelloWorld接口
 * - 测试指定ES连接的Ping操作
 * - 返回ES连接的状态和响应信息
 * 
 * @param data 请求参数对象
 * @param data.es_connect ES连接ID，通常通过sdk.GetSelectEsConnID()获取
 * @param data.text 可选的文本参数，用于扩展功能
 * @returns Promise<ApiResponse> 返回API响应，包含code、msg、data字段
 * 
 * 使用示例：
 * ```typescript
 * import { HelloAction } from '@/api/helloworld'
 * import { sdk } from '@elasticview/plugin-sdk'
 * 
 * const result = await HelloAction({
 *   es_connect: sdk.GetSelectEsConnID()
 * })
 * ```
 */
export function HelloAction(data: any) {
  return request({
    url: `/api/HelloWorld`,  // 请求路径：对应后端router中注册的HelloWorld接口
    method: 'post',          // HTTP方法：使用POST方法发送请求
    data                     // 请求数据：传递给后端的参数对象
  })
}
