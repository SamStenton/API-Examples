#include <iostream>
#include <cpr/cpr.h>

int main(int argc, char** argv) {
    cpr::Response r = cpr::Get(cpr::Url{ "https://engineering-task.elancoapps.com/api/applications" });
    /*
    r.status_code;                  // 200
    r.header["content-type"];       // application/json; charset=utf-8
    r.text;                         // JSON text string
    */
    std::cout << r.text;
    system("pause");
    return 0;
}