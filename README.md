

# Amber

This project was generated using [Nx](https://nx.dev).

## Steps

- Install Visual studio - build tools
- Generate Project (select apps and no cloud)
  - `npm exec create-nx-workspace amber --preset=empty --cli=nx`
  - `cd amber`
- Install nx-go plugin
  - `npm install -D @nx-go/nx-go`
- Create Apps
  - `nx g @nx-go/nx-go:app as`
  - `nx g @nx-go/nx-go:app ps`
- Create Libs
  - `nx g @nx-go/nx-go:lib crypt`
- Add application and libraries code
- Check dependency graph
  - `nx graph` 
- Ready to go
  - `nx build as`
  - `nx lint as`
  - `nx test as`
  - `nx serve as`
- Check dependencies affected by change and test only afftected components
  - `nx affected:dep-graph`
  - `nx affected:test --parallel`

## References
https://github.com/nx-go/nx-go
https://nx.dev/core-tutorial/01-create-blog
