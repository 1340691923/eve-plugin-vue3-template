// ElasticView插件数据库测试 API接口文件 - 定义与插件数据存储相关的HTTP请求
// 演示ElasticView插件的数据库CRUD操作能力

import {request} from '@elasticview/plugin-sdk' // 导入ElasticView插件SDK的HTTP请求函数

/**
 * DbInsert API接口 - 向插件数据库插入新记录
 * 
 * 功能说明：
 * - 调用插件后端的数据库插入接口
 * - 将键值对数据保存到插件专用数据库
 * - 支持SQLite和MySQL数据库
 * 
 * @param data 请求参数对象
 * @param data.key 要插入的数据键名
 * @param data.value 要插入的数据值
 * @returns Promise<ApiResponse> 返回操作结果，包含code、msg字段
 */
export function DbInsert(data: any) {
    return request({
        url: `/api/DbInsert`,    // 请求路径：对应后端router中注册的DbInsert接口
        method: 'post',          // HTTP方法：使用POST方法发送请求
        data                     // 请求数据：包含key和value的对象
    })
}

/**
 * CleanAction API接口 - 清空插件数据库中的所有数据
 * 
 * 功能说明：
 * - 调用插件后端的数据库清空接口
 * - 删除插件数据库中的所有记录
 * - 用于重置插件数据状态
 * 
 * @param data 请求参数对象（通常为空对象）
 * @returns Promise<ApiResponse> 返回操作结果，包含code、msg字段
 */
export function CleanAction(data: any) {
    return request({
        url: `/api/DbDelete`,    // 请求路径：对应后端router中注册的DbDelete接口
        method: 'post',          // HTTP方法：使用POST方法发送请求
        data                     // 请求数据：清空操作通常不需要参数
    })
}

/**
 * ListAction API接口 - 分页查询插件数据库中的数据
 * 
 * 功能说明：
 * - 调用插件后端的数据库查询接口
 * - 支持分页查询，返回数据列表和总数
 * - 用于展示插件存储的数据内容
 * 
 * @param data 请求参数对象
 * @param data.page 页码，从1开始
 * @param data.limit 每页记录数
 * @returns Promise<ApiResponse> 返回查询结果，data字段包含list数组和count总数
 */
export function ListAction(data: any) {
    return request({
        url: `/api/DbSearch`,    // 请求路径：对应后端router中注册的DbSearch接口
        method: 'post',          // HTTP方法：使用POST方法发送请求
        data                     // 请求数据：包含分页参数的对象
    })
}
