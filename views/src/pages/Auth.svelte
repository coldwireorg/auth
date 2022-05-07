<script>
  import { location, push } from 'svelte-spa-router'
  import oxyd from '../libs/oxyd'

  import Link from '../components/Link.svelte'
  import Field from '../components/Field.svelte'
  import Button from '../components/Button.svelte'

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

      let keys = await oxyd.generate(password)
      req.public_key = oxyd.utils.u8ToBase64(keys.public_key)
      req.private_key = oxyd.utils.u8ToBase64(keys.private_key)
    }

    const res = await fetch(`/api/auth/${$location === '/sign-up' ? 'register' : 'login'}${challenge != null ? '?login_challenge='+challenge : ''}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(req)
    })

    const data = await res.json()

    if (data.status == "SUCCESS") {
      if (data.content.redirect_url) {
        window.location.replace(data.content.redirect_url)
      } else {
        push("/user/hub")
      }
    } else {
      error = data.message
    }
  }
</script>

<div class="background"></div>
<a class="art" href="https://aenamiart.artstation.com/" target="_blank">Art by Alena Aenami</a>

<div class="auth">
  <div class="form">
    <div class="top">
      <img src="/icons/lock.svg" alt="">
      <p>Make sure your connection is safe</p> 
    </div>
    
    <h4 class="title">{$location === '/sign-up' ? 'Sign-up' : 'Sign-in'}</h4>

    <div class="fields">
      {#if error != ""}
        <span class="error">{error}</span>
      {/if}
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
  .background {
    background: url(/img/background.jpg) no-repeat center;
    background-size: cover;

    position: absolute;
    left: 0; right: 0;
    top: 0; bottom: 0;

    z-index: -1;
  }

  .art {
    position: absolute;
    bottom: 32px;
    left: 32px;
    background-color: var(--complementary-gray-3);
    padding: 8px 16px 8px 16px;
    border-radius: 8px;
    z-index: 50;
  }

  .auth {
    position: absolute;
    left: 0; bottom: 0; top: 0;
    width: 100%;

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

  .form .fields .error {
    background-color: var(--secondary-red);
    padding: 8px;
    border-radius: 8px;
  }

  .form .question p {
    font-size: 10px;
    color: var(--complementary-white-50);
  }
</style>