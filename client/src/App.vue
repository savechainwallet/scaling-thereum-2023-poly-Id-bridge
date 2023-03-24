<template lang="pug">
div(id="app")
    b-navbar(toggleable="lg" type="dark" variant="info")
      b-navbar-brand(href="#") PolyIDBridge
      b-navbar-toggle(target="nav-collapse")
      b-collapse(id="nav-collapse" is-nav)
        b-navbar-nav
          b-nav-item(to="/") Sign In
          b-nav-item(to="/profile") Profile
    
    router-view
</template>

<style scoped>
 .navbar.navbar-dark.bg-info{
    background-color: #0d9263!important;
  
 } 

</style>
