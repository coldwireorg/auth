<script>
	import Router, { location, push } from 'svelte-spa-router'
  import {wrap} from 'svelte-spa-router/wrap'

  import User from './pages/User.svelte'
  import Auth from './pages/Auth.svelte'

  const isAuth = async () => {
    const res = await fetch(`/api/user/info`)
		if (res.status === 200) {
      return true
    } else {
      return false
    }
  }

  const routes = {
    '/': wrap({
      component: User,
      conditions: [
        async (detail) => {
          if (await isAuth()) {
            push("/user/hub")
          } else {
            push("/sign-in")
          }
        }
      ]
    }),
    '/user/:page?': wrap({
      component: User,
      conditions: [
        async (detail) => {
          if (await isAuth()) {
            return true
          } else {
            push("/sign-in")
          }
        }
      ]
    }),
    '/sign-in': wrap({
      component: Auth,
      conditions: [
        async (detail) => {
          if (await isAuth()) {
            push("/user/hub")
          } else {
            return true
          }
        }
      ]
    }),
    '/sign-up': wrap({
      component: Auth,
      conditions: [
        async (detail) => {
          if (await isAuth()) {
            push("/user/hub")
          } else {
            return true
          }
        }
      ]
    })
  }
</script>

<main>
  <a class="logo" href="https://coldwire.org">
    <div />
  </a>

  <Router {routes} />
</main>

<style>
  .logo {
    z-index: 5000;
    position: absolute;
    left: 32px;
    top: 24px;
  }

  .logo div {
    background: url(/img/logo.svg);
    background-size: cover;

    height: 16px;
    width: 27.65px;
  }
</style>