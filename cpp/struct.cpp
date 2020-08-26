#include <cstdio>

using namespace std;

struct Employee {
int id;
const char * name;
const char * role;
};

int main() {
Employee joe = {42, "Joe", "Boss"};
Employee * e = &joe;

printf("%s is the %s and has id %d\n",
e->name, e->role, e->id);

return 0
}
