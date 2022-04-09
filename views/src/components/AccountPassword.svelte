<script>
  import Card from '../components/Card.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte';
  import Button from '../components/Button.svelte';
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <Text type="h3">
    Change your password
  </Text>
  <br>
  <Field type="password" placeholder="Enter your new password" />
  <Field type="password" placeholder="Repeat your new password" />
  <Field type="password" placeholder="Enter your current password" />
  <br>
  <Button>
    Change password
  </Button>
</Card>