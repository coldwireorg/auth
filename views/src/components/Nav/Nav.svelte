<script>
  import AccountCard from './NavAccount.svelte'
  import AccountDonate from './NavDonate.svelte'

  export let usr
</script>

<div class="nav">
  <div class="nav-content">
    <div id="top">
      <a class="link" href="/#/user/hub"><img src="/icons/services.svg" alt=""> Hub</a>
      <!-- <a class="link" href="/#/user/blog"><img src="/icons/info.svg" alt=""> Blog</a> -->
      {#if usr.role == "admin"}
        <a class="link" href="/#/user/admin"><img src="/icons/admin.svg" alt=""> Admin</a>
      {/if}
    </div>
    <div id="bottom">
      <AccountCard usrdta={usr} />
      <AccountDonate />
    </div>
  </div>
</div>

<style>
  .nav {
    width: 254px;
    padding: 0;
    padding-top: 94px;
    margin: 0;

    display: flex;
    flex-direction: column;
    gap: 16px;

    background-color: var(--complementary-gray-3);
  }

  .nav .nav-content {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 16px;
    height: 100%;
    margin: 32px;
  }

  .nav .nav-content #top {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .nav .nav-content #top .link {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px;
    border-radius: 8px;
  }

  .nav .nav-content #top .link:hover {
    background: rgba(240, 246, 252, 0.05);
  }

  .nav .nav-content #top .link img {
    height: 16px;
  }

  .nav .nav-content #bottom {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
</style>