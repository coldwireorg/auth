<script>
  import { onMount } from 'svelte';
  import { push, location } from 'svelte-spa-router';
  import { fade } from 'svelte/transition'

  import Nav from '../components/Nav/Nav.svelte';
  import Hub from '../components/Hub/Hub.svelte';

  if ($location == "/") {
    push("/hub")
  }

  let userData = {
    username: '',
    public_key: '',
    role: ''
  }

  onMount(async () => {
		const res = await fetch(`/api/user/info`)
		const r = await res.json()
    
    userData.username = r.content.username
    userData.role = r.content.group
    userData.public_key = r.content.public_key
	})
</script>

{#if userData.username != ''}
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
{/if}

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