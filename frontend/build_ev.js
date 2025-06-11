// ElasticView插件构建后处理脚本 - 修复生产环境下的动态路径问题
// 此脚本在Vite构建完成后运行，处理微前端环境下的路径注入问题

import fs from 'fs/promises'; // 导入异步文件系统操作模块
import path from 'path';      // 导入路径处理模块
import url from 'url';        // 导入URL处理模块

// 获取当前脚本的文件路径和目录路径 (ES模块环境下的__filename和__dirname替代方案)
const __filename = url.fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// 构建输出的HTML文件路径 - 指向Vite构建产物中的index.html
const filePath = path.resolve(__dirname, 'dist', 'index.html');

// updateHtml异步函数 - 修复HTML文件中的动态基础路径引用
// ElasticView插件系统使用qiankun微前端架构，需要在运行时动态注入基础路径
async function updateHtml() {
    try {
        // 读取构建产物中的HTML文件内容
        const content = await fs.readFile(filePath, 'utf8');

        // 使用正则表达式替换所有的静态路径引用为动态路径引用
        // 将形如 '/__dynamic_base__ 的静态引用替换为 window.__dynamic_base__+' 的动态引用
        // 这样ElasticView插件系统就可以在运行时注入正确的基础路径
        const modifiedContent = content.replaceAll("'/__dynamic_base__", "window.__dynamic_base__+'");

        // 将修改后的内容写回HTML文件
        await fs.writeFile(filePath, modifiedContent, 'utf8');
        console.log('dist/index.html 文件内容已更新。'); // 输出成功信息
    } catch (err) {
        console.error('处理文件时出错:', err); // 输出错误信息
    }
}

// 执行HTML文件更新操作
// 这是构建流程的最后一步，确保插件可以正确在ElasticView系统中运行
await updateHtml()
