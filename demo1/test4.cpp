#include<stack>
#include<vector>
#include<algorithm>
#include<iostream>
using namespace std;


struct Queen
{

    unsigned int x, y;
    Queen(unsigned int xx = 0, unsigned int yy = 0) :x(xx), y(yy){};
    //重载相等与不等运算符
    bool operator==(const Queen& qa)const
    {
        return(x == qa.x || y == qa.y || (x + y) == (qa.x + qa.y) || (x - y) == (qa.x - qa.y));
    }
    bool operator!=(Queen const& qa)const
    {
        return !(*this == qa);
    }
    //重载赋值运算符
    Queen& operator=(Queen const& qa)
    {
        x = qa.x;
        y = qa.y;
        return *this;
    }
};

void myN_Queens(unsigned int N)
{
    vector<Queen> result;
    Queen q(0, 0);
    unsigned int row = 0;//标记当前在操作的行，当操作到了最后一行的后面，就代表可以结束了
    while (row < N)
    {
        while ((q.y < N) && (result.end() != find(result.begin(), result.end(), q)))
            //直到在本行找到一个位置或者到了行尾,不然就一直往右移试探本行的下一位置
        {
            q.y++;
        }
        if (N > q.y)//如果没到行尾，也就是在本行找到了一个合适的位置
        {
            result.insert(result.end(), q);//把此时的位置入栈
            q.x++; q.y = 0;//转到下一行的行首
            row++;
        }
        else//到了行尾
        {
            q = result[result.size() - 1];//把上一行的坐标再拿出来
            result.erase(result.end() - 1);//删除之，其实就是出栈的过程
            row--;
            q.y++;//本行当前位置不合适，继续往后找一个合适的位置
        }

    }

    for (auto i : result)
        cout << (i).x << " " << (i).y << endl;
    cout << "ok" << endl;
}

int main() {
    int n;
    cin>>n;
    myN_Queens(n);
}