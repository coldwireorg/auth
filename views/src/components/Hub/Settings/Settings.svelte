<script>
  import SettingsDelete from "./SettingsDelete.svelte";
  import SettingsPassword from "./SettingsPassword.svelte";
</script>

<div class="settings">
  <SettingsPassword />
  <SettingsDelete />
</div>

<style>
  .settings {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 16px;
    max-width: 528px;
  }
</style>