# CLAUDE.md

此檔案為 Claude Code (claude.ai/code) 在本專案中工作時提供指引。

## 專案概述

單檔待辦事項應用程式，使用純 HTML/CSS/JavaScript 建構。無建構工具、無依賴、無框架。

## 執行方式

在瀏覽器中直接開啟 `index.html`（macOS 上使用 `open index.html`）。

## 架構

所有程式碼都在 `index.html` 中：

- **CSS** 在 `<style>` 內 — 使用 Flexbox 佈局，最大寬度 520px 置中容器
- **JS** 在 `<script>` 內 — 原生 DOM 操作，搭配 `localStorage` 做資料持久化
- 介面語言為繁體中文（zh-TW）

資料模型：`todos` 陣列，元素結構為 `{ id: number, text: string, done: boolean }`，儲存在 `localStorage` 的 `'todos'` 鍵值下。

核心函式：`addTodo`、`toggleTodo`、`deleteTodo`、`editTodo`（雙擊行內編輯，透過 contentEditable）、`setFilter`、`clearDone`、`render`（每次狀態變更時重新渲染整個列表）。
