#include <vector>
#include <iostream>
#include "sort.h"
#include <algorithm>

using namespace std;

int cpp() {
    vector<int> v;
    for (int i =3 ; i >= 0; i-- ){
        v.push_back(i);
        cout << i << endl;
    }
    sort(v.begin(), v.end());
    cout << "After sort";
    for (int i = 0 ; i < v.size(); i++ ){
        cout << i << endl;
    }
    return v[0];
}
