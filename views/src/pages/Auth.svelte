<script>
  import { location, push } from 'svelte-spa-router'
  import { fade } from 'svelte/transition'

  import Link from '../components/Link.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';

  const urlParams = new URLSearchParams(window.location.search);

  let username = ""
  let password = ""
  let repassword = ""

  let error = ""

  const auth = async () => {
    const challenge = urlParams.get("login_challenge")

    let req = {
      username: username,
      password: password,
      private_key: "",
      public_key: ""
    }

    if ($location === '/sign-up') {
      if (password != repassword) {
        error = "passwords does not match"
        return
      }

      let keys = oxyd.generate(password)

      req.public_key = keys.publicKey
      req.private_key = keys.privateKey
    }

    const res = await fetch(`/api/auth/${$location === '/sign-up' ? 'register' : 'login'}${challenge != '' ? '?login_challenge='+challenge : ''}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(req)
    })

    const data = await res.json()

    console.log(data)

    if (data.status == "SUCCESS") {
      if (data.content.redirect_url) {
        window.location.replace(data.content.redirect_url)
      } else {
        push("/user")
      }
    }
  }
</script>

<div class="auth" in:fade={{duration: 300}} out:fade={{duration: 300}}>
  <div class="form">
    <div class="top">
      <img src="/icons/lock.svg" alt="">
      <p>Make sure your connection is safe</p> 
    </div>
    
    <h4 class="title">{$location === '/sign-up' ? 'Sign-up' : 'Sign-in'}</h4>

    <div class="fields">
      <Field type="text" placeholder="Username" bind:value={username} />
      <Field type="password" placeholder="password" bind:value={password} />
      {#if $location === '/sign-up'}
        <Field type="password" placeholder="Repeat password" bind:value={repassword} />
      {/if}
      <Button on:click={auth}>
        {$location === '/sign-up' ? 'Sign-up' : 'Sign-in'}
      </Button>
      <div class="question">
        {#if $location === '/sign-in'}
          <p>You don't have an account yet? <Link link="/sign-up">Create one</Link>!</p>
        {:else}
          <p>Already have an account?? <Link link="/sign-in">Sign in</Link>!</p>
        {/if}
      </div>
    </div>
  </div>  
</div>

<style>
  .auth {
    position: absolute;
    left: 0; bottom: 0; top: 0;
    width: 50%;

    display: flex;
    justify-content: center;
    align-items: center;
  }

  .form {
    padding: 32px;
    border-radius: 8px;
    width: 212px;
    background-color: var(--complementary-gray-3);
  }

  .form .top {
    margin: -32px;
    margin-bottom: 0px;
    border-top-left-radius: 8px;
    border-top-right-radius: 8px;
    padding-left: 32px;
    padding-right: 32px;

    height: 46px;
    width: 100%;

    background-color: var(--primary-green-25);
    color: var(--primary-green);

    display: flex;
    justify-content: space-around;
    align-items: center;
  }

  .form .top p {
    font-size: 10px;
  }

  .form .title {
    margin-top: 32px;
    margin-bottom: 32px;
  }

  .form .fields {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .form .question p {
    font-size: 10px;
    color: var(--complementary-white-50);
  }
</style>