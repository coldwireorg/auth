<script>
  import Card from '../components/Card.svelte'
  import Profile from '../components/Profile.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';
</script>

<div class="user">
  <dir class="account">
    <Text type="h2">Account</Text>
    <Card style="display: flex; flex-direction: column; gap:8px;">
      <div class="user-name">
        <Profile username="monoko" />
        <Text type="h2">Monoko</Text>
      </div>
      <br>
      <div class="account-info-item">
        <img src="/icons/key.svg" alt="">
        <Text type="h3" color="var(--complementary-white-50)">EgNdEt1ue2DtqR9/qCckf7tfifNJ+vHd...</Text>
      </div>
      <div class="account-info-item">
        <img src="/icons/eye.svg" alt="">
        <Text type="h3" color="var(--complementary-white-50)">Acces level 1</Text>
      </div>
      <div class="account-info-item">
        <img src="/icons/server.svg" alt="">
        <Text type="h3" color="var(--complementary-white-50)">coldwire.org</Text>
      </div>
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

  </div>
</div>

<style>
  .account {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;

    width: 316px;
    padding: 32px;
    padding-top: 94px;
    margin: 0;

    display: flex;
    flex-direction: column;
    gap: 32px;

    background-color: var(--complementary-gray-2);
  }

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