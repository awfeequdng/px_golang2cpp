# px_golang2cpp
convert go to cpp
- convert map to std::unordered_map
- convert array to vector
- convert 'var test bool' to 'bool test{}'
- convert 'var test = 2' to 'auto test = 2;'
- convert 'var test []string' to 'vector<string> test{};'
- convert 'var test map[int]string' to std::unordered_map<int,string> test{};
