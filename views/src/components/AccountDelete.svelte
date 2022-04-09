<script>
  import Card from '../components/Card.svelte'
  import Text from '../components/Text.svelte'
  import Field from '../components/Field.svelte'
  import Button from '../components/Button.svelte'

  let password = ""
  let error = ""

  const del = () => {
    if (password == "") {
      error = "please fill the password field"
      return
    }

    let conf = confirm("Do you really want to delete your account?")
    if (!conf) return

    fetch("/api/user", {
      method: 'DELETE',
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        password: password,
      })
    })
    .then(req => req.json())
    .then(data => {
      if (data.status == "SUCCESS") {
        location.replace("/api/auth/logout")
      } else {
        error = data.message
      }
    })
  }
</script>

<Card style="display: flex; flex-direction: column; gap:8px;">
  <Text type="h3">
    Delete your account
  </Text>
  <br>
  {#if error != ""}
    <span class="error">{error}</span>
  {/if}
  <Field type="password" placeholder="Enter password" bind:value={password} />
  <br>
  <Button color="red" on:click={del}>
    Delete account
  </Button>
</Card>

<style>
  .error {
    background-color: var(--secondary-red);
    padding: 8px;
    border-radius: 8px;
  }
</style>