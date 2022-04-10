<script>
  import Card from '../Card.svelte'
  import Profile from '../Profile.svelte'
  import Text from '../Text.svelte'

  export let usrdta

  const logout = () => [
    fetch("/api/auth/logout")
    .then(() => window.location.reload())
  ]
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <div class="user">
    <div class="info">
      <Profile username={usrdta.username} />
      <Text type="h3">{usrdta.username}</Text>
    </div>
    <span class="logout" on:click={logout}>
  </div>
  <br>
  <div class="account-info-item">
    <img src="/icons/key.svg" alt="">
    <Text type="h3" color="var(--complementary-white-50)">{usrdta.public_key.slice(0, 15)}...</Text>
  </div>
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

  .user {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .user .logout {
    height: 24px;
    width: 24px;
    border-radius: 25px;
    background: url(/icons/logout.svg) no-repeat center;
    background-size: 70%;
    transition: 0.3s;
    cursor: pointer;
  }

  .user .logout:hover {
    background-color: var(--complementary-white-5);
  }

  .user .info {
    display: flex;
    gap: 16px;
    align-items: center;
  }
</style>