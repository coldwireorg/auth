<script>
  import Card from '../components/Card.svelte'
  import Profile from '../components/Profile.svelte'
  import Text from '../components/Text.svelte'
  import Button from '../components/Button.svelte';

  export let usrdta

  const logout = () => [
    fetch("/api/auth/logout")
    .then(() => window.location.reload())
  ]
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <div class="user-name">
    <Profile username={usrdta.username} />
    <Text type="h2">{usrdta.username}</Text>
  </div>
  <br>
  <div class="account-info-item">
    <img src="/icons/key.svg" alt="">
    <Text type="h3" color="var(--complementary-white-50)">{usrdta.public_key.slice(0, 30)}...</Text>
  </div>
  <div class="account-info-item">
    <img src="/icons/eye.svg" alt="">
    <Text type="h3" color="var(--complementary-white-50)">Acces level 1</Text>
  </div>
  <div class="account-info-item">
    <img src="/icons/server.svg" alt="">
    <Text type="h3" color="var(--complementary-white-50)">coldwire.org</Text>
  </div>
  <br>
  <Button color="red" on:click={logout}>
    Logout
  </Button>
</Card>

<style>
  .account-info-item {
    display: flex;
    gap: 8px;
    align-items: center;
  }

  .account-info-item img {
    width: 12px;
  }

  .user-name {
    display: flex;
    gap: 16px;
    align-items: center;
  }
</style>