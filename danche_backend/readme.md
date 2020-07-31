node_modules  不应该被跟踪，太大了
npm i package.json 

ls 查看当前目录下所有文件 
ls -al + 文件名   查看该文件下的所有文件
cd + 文件名/  进入该文件目录
pwd 展示当前文件的绝对物理路径
du -hs  +  文件名  查看该文件的大小
cd .. 回到上级目录
cat + 文件  输出该文件


把项目clone到电脑想要clone的地方
clone之后
yarn  安装依赖
git pull origin master 更新clone的文件


1. 删除node_modules
 git rm

2. .gitignore
3. 提交删除