<script>
  import { fade } from 'svelte/transition'
  import jwt_decode from "jwt-decode";

  import AccountCard from '../components/AccountCard.svelte'
  import AccountPassword from '../components/AccountPassword.svelte'
  import AccountDelete from '../components/AccountDelete.svelte'
  import AccountDonate from '../components/AccountDonate.svelte'
  import Hub from '../components/Hub.svelte';

  import Text from '../components/Text.svelte'

  let userData = jwt_decode(document.cookie.split("=")[1])
</script>

<div class="user" in:fade={{duration: 300}} out:fade={{duration: 300}}>
  <dir class="account">
    <Text type="h2">Account</Text>
    <div class="widgets">
      <div id="top">
        <AccountCard usrdta={userData} />
        <AccountPassword />
        <AccountDelete />
      </div>
      <AccountDonate />
    </div>
  </dir>
  <Hub />
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

  .user .account {
    width: 254px;
    padding: 32px;
    padding-top: 94px;
    margin: 0;

    display: flex;
    flex-direction: column;
    gap: 16px;

    background-color: var(--complementary-gray-3);
  }

  .user .account .widgets {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 16px;
    height: 100%;
  }

  .user .account .widgets #top {
    display: flex;
    flex-direction: column;
    gap: 32px;
  }
</style>