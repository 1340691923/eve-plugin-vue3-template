// ElasticView插件英文语言包 - 提供插件界面的英文文本资源
// 根据ElasticView插件开发规范，支持国际化多语言切换

/**
 * 英文语言包配置对象
 * 
 * 使用方式：
 * - 在Vue模板中：{{ $t('click me') }}
 * - 在JavaScript中：import { useI18n } from 'vue-i18n'; const { t } = useI18n(); t('click me')
 * 
 * 注意事项：
 * - 键名应与中文语言包保持一致
 * - 建议使用嵌套对象组织复杂的文本结构
 * - 支持参数化文本，如：'Hello {name}'
 */
export default {
    // 通用操作文本
    'click me': 'click me',        // 点击按钮文本
    '刷新': 'Refresh',             // 刷新操作
    '清空': 'Clear',               // 清空操作
    '新增': 'Add',                 // 新增操作
    '保存': 'Save',                // 保存操作
    '操作': 'Actions',             // 操作列标题
    
    // 页面标题和描述
    'HelloWorld': 'HelloWorld',    // HelloWorld页面标题
    'Database Test': 'Database Test', // 数据库测试页面标题
    'ES Connection Test': 'ElasticSearch Connection Test', // ES连接测试
    
    // 表单和数据
    '键': 'Key',                   // 数据键字段
    '值': 'Value',                 // 数据值字段
    '数据库操作 demo': 'Database Operations Demo', // 数据库操作演示
    
    // 消息提示
    '操作成功': 'Operation Successful',    // 成功提示
    '操作失败': 'Operation Failed',        // 失败提示
    '请求错误': 'Request Error',           // 请求错误提示
};
