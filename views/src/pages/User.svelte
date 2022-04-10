<script>
  import { push, location } from 'svelte-spa-router'
  import { fade } from 'svelte/transition'
  import jwt_decode from "jwt-decode";


  import Nav from '../components/Nav/Nav.svelte';
  import Hub from '../components/Hub/Hub.svelte';

  if ($location == "/") {
    push("/hub")
  }

  let userData = jwt_decode(document.cookie.split("=")[1])
</script>

<div class="user" in:fade={{duration: 300}} out:fade={{duration: 300}}>
  <Nav usr={userData} />
  {#if $location == "/user/hub"}
    <Hub />
  {:else if $location == "/user/blog"}
    blog
  {:else if $location == "/user/admin"}
    admin
  {/if}
</div>

<style>
  .user {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    display: flex;
  }
</style>