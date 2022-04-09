<script>
  import { fade } from 'svelte/transition'
  import jwt_decode from "jwt-decode";

  import Card from '../components/Card.svelte'
  import Profile from '../components/Profile.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';
  import Tag from '../components/Tag.svelte';

  let userData = jwt_decode(document.cookie.split("=")[1])

  const logout = () => [
    fetch("/api/auth/logout")
    .then(() => window.location.reload())
  ]
</script>

<div class="user" in:fade={{duration: 300}} out:fade={{duration: 300}}>
  <dir class="account">
    <Text type="h2">Account</Text>
    <Card style="display: flex; flex-direction: column; gap:8px;">
      <div class="user-name">
        <Profile username={userData.username} />
        <Text type="h2">{userData.username}</Text>
      </div>
      <br>
      <div class="account-info-item">
        <img src="/icons/key.svg" alt="">
        <Text type="h3" color="var(--complementary-white-50)">{userData.public_key.slice(0, 30)}...</Text>
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
    <Card style="display: flex; flex-direction: column; gap:8px;">
      <Text type="h2">Change your password</Text>
      <br>
      <Field type="password" placeholder="Enter your new password" />
      <Field type="password" placeholder="Repeat your new password" />
      <Field type="password" placeholder="Enter your current password" />
      <br>
      <Button>
        Change password
      </Button>
    </Card>
    <Card type="green-25" style="display: flex; flex-direction: column; gap:8px;">
      <Text type="h3">Help us to keep our services alive!</Text>
      <Button color="green">Donate</Button>
    </Card>
  </dir>
  <div class="hub">
    <Text type="h2">Services</Text>
    <div class="content">
      <div class="category">
        <div class="header">
          <img src="/icons/flame.svg" alt="">
          <Text type="h3">Organizing</Text>
        </div>
  
        <div class="services">
          <div class="service">
            <Text type="h2">Matrix</Text>
            <Text type="p" color="var(--complementary-white-50)">Discuss safely</Text>
          </div>
          <div class="service">
            <Text type="h2">Bloc</Text>
            <Text type="p" color="var(--complementary-white-50)">Store and share your files</Text>
          </div>
          <div class="service">
            <div class="header">
              <Text type="h2">Riot</Text>
              <Tag text="Soon" background="green" />
            </div>
            <Text type="p" color="var(--complementary-white-50)">Organize protests</Text>
          </div>
        </div>
      </div>

      <div class="category">
        <div class="header">
          <img src="/icons/hosting.svg" alt="">
          <Text type="h3">Hosting</Text>
        </div>
  
        <div class="services">
          <div class="service">
            <div class="header">
              <Text type="h2">VPS</Text>
              <Tag text="On request" background="purple" />
              <Tag text="Soon" background="green" />
            </div>
            <Text type="p" color="var(--complementary-white-50)">Hidden VPS services</Text>
          </div>
        </div>
      </div>
    </div>
  </div>
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

  .user .account-info-item {
    display: flex;
    gap: 8px;
    align-items: center;
  }

  .user .account-info-item img {
    width: 12px;
  }

  .user .user-name {
    display: flex;
    gap: 16px;
    align-items: center;
  }

  .user .hub {
    margin: 32px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .user .hub .content {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .user .hub .content .category {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 16px;
    gap: 16px;

    background: #161B22;
    border-radius: 8px;

    width: 496px;
  }

  .user .hub .content .category .header {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .user .hub .content .category .header img {
    height: 16px;
  }

  .user .hub .content .category .services {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    width: 100%;
  }

  .user .hub .content .category .services .service {
    background: #30363D;
    border-radius: 8px;
    padding: 16px;
    width: calc(50% - 36px);
  }

  .user .hub .content .category .service .header {
    
  }
</style>