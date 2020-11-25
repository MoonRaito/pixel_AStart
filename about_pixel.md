### pixel 部分属性

R Rect 矩形

Matrix 矩阵

向量 位置、运动（平移）、速度、加速度等。
Vector, pixel.Vec. Positions, movements (translations), velocities, accelerations, and so on.

矩形 主要是相框（精灵的一部分）和边界。
Rectangle, pixel.Rect. Mainly picture frames (portions for sprites) and bounds.
自理解：它具有Min和Max组件。Min是矩形左下角Max的位置，是矩形右上角的位置。矩形的边始终与X和Y轴平行。


矩阵 各种线性变换：移动、旋转、缩放。
Matrix, pixel.Matrix. All kinds of linear transformations: movements, rotations, scaling.


https://github.com/faiface/pixel/wiki/Moving,-scaling-and-rotating-with-Matrix