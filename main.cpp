#include <bits/stdc++.h>

using namespace std;

int maxSatisfied(vector<int> customers, vector<int>grumpy, int minutes) {
    int totalCustomers = 0;
    int totalUnsatisfiedCustomers = 0;
    for (int i = 0; i < customers.size(); i++) {
        totalCustomers += customers[i];
        customers[i] *= grumpy[i];
        totalUnsatisfiedCustomers += customers[i];
    }

    int left = 0;
    int right = left + minutes - 1;

    int satisfiedCustomersInRage = 0;
    for (int i = left; i <= right; i++)
    {
        satisfiedCustomersInRage += customers[i];
    }

    int maxSatisfiedCustomers = satisfiedCustomersInRage;

    while (right < customers.size() -1) {
        satisfiedCustomersInRage -= customers[left];
        left++;
        satisfiedCustomersInRage += customers[right + 1];
        right++;
        maxSatisfiedCustomers = max(maxSatisfiedCustomers, satisfiedCustomersInRage);
    }

    return totalCustomers - totalUnsatisfiedCustomers + maxSatisfiedCustomers;
}

int main() {
    vector<int> customers {1,0,1,2,1,1,7,5};
    vector<int> grumpy {0,1,0,1,0,1,0,1};
    int minutes = 3;

    // vector<int> customers {10,1,7};
    // vector<int> grumpy {0,0,0};
    // int minutes = 2;

    cout << maxSatisfied(customers, grumpy, minutes) << endl; // 16 (1 + 1 + 1 + 7 + 5

    return 0;
}