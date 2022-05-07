<script>
	import Router, { location, push } from 'svelte-spa-router'
  import {wrap} from 'svelte-spa-router/wrap'

  import User from './pages/User.svelte'
  import Auth from './pages/Auth.svelte'

  const isAuth = () => {
    if (document.cookie.split("=")[1]) {
      return true
    }

    return false
  }

  const routes = {
    '/': wrap({
      component: User,
      conditions: [
        (detail) => {
          if (isAuth()) {
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
        (detail) => {
          if (isAuth()) {
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
        (detail) => {
          if (isAuth()) {
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
        (detail) => {
          if (isAuth()) {
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