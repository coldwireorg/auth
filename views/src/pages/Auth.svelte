<script>
  import { location } from 'svelte-spa-router'

  import Link from '../components/Link.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';
</script>

<div class="auth">
  <div class="form">
    <div class="top">
      <img src="/icons/lock.svg" alt="">
      <p>Make sure your connection is safe</p> 
    </div>
    
    <h4 class="title">{$location === '/sign-up' ? 'Sign-up' : 'Sign-in'}</h4>

    <div class="fields">
      <Field type="text" placeholder="Username" />
      <Field type="password" placeholder="password" />
      {#if $location === '/sign-up'}
        <Field type="password" placeholder="Repeat password" />
      {/if}
      <Button link="#">
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