<script>
  import { fade } from 'svelte/transition'
  import jwt_decode from "jwt-decode";

  import AccountCard from '../components/AccountCard.svelte'
  import AccountPassword from '../components/AccountPassword.svelte'
  import AccountDonate from '../components/AccountDonate.svelte'
  import Hub from '../components/Hub.svelte';

  import Text from '../components/Text.svelte'

  let userData = jwt_decode(document.cookie.split("=")[1])

  let categories = [
    {
      name: "Organizing",
      icon: "flame",
      services: [
        {
          name: "Matrix",
          desc: "Chat safely",
        },
        {
          name: "Bloc",
          desc: "Securely store your files",
          tags: [
            {text: "soon", color: "green"}
          ]
        },
      ]
    },
    {
      name: "Hosting",
      icon: "hosting",
      services: [
        {
          name: "VPS",
          desc: "Run your own server anonymously",
          tags: [
            {text: "soon", color: "green"},
            {text: "On request", color: "purple"}
          ]
        },
      ]
    }
  ]
</script>

<div class="user" in:fade={{duration: 300}} out:fade={{duration: 300}}>
  <dir class="account">
    <Text type="h2">Account</Text>
    <AccountCard usrdta={userData} />
    <AccountPassword />
    <AccountDonate />
  </dir>
  <Hub categories={categories} />
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
    width: 316px;
    padding: 32px;
    padding-top: 94px;
    margin: 0;

    display: flex;
    flex-direction: column;
    gap: 32px;

    background-color: var(--complementary-gray-2);
  }
</style>