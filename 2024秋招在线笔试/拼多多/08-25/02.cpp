#include <iostream>
#include <string>

int main() {
    int n;
    std::string name;
    std::string age;

    std::cin >> n;
    std::cin >> name >> age;

    // 处理输入的数据
    std::cout << "Name: " << name << ", Age: " << age << std::endl;

    return 0;
}
