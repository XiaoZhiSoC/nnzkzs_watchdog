# nnzkzs_watchdog
愿风指引你的道路，愿星辰照亮你前进的方向。

## 用途

实时监视南宁中考招生排名情况和差距。截至2023.07.13，南宁中考招生网站上难以查看未入围学校的具体排名情况以及差距。这个工具可以更好的“捡漏”。

## 背景

中考B，捡漏写的一个实时监测系。虽然最后也没去成那个高中（最后20s软件输出正数我就知道被踢了）不过也还好吧，掐在最后改报了。

## 使用

- 在网站上登录账号后截取 Cookie 以及 X-Xsrf-Token 参数填到对应位置（两个参数都在Headers部分）。
- 直接编译运行，也可以使用IDE。
- 差距部分是负数就是入围（差距的绝对值是后面排有多少人）差距部分是正数就说不入围（差距的绝对值是距离入围最底线还有多少人）

## 须知

- 这是一个临时写出来的东西，开源纪念高中报名这天。
- 仅用于南宁中考招生网。

> 该项目不会有后续更新，也无法更新。
