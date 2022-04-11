<script>
  import HubPanel from './HubPanel.svelte'
  import Services from './Services/Services.svelte'
  import Settings from './Settings/Settings.svelte'
</script>

<div class="hub">
  <HubPanel title="Services">
    <Services />
  </HubPanel>
  <HubPanel title="Settings">
    <Settings />
  </HubPanel>
</div>

<style>
  .hub {
    margin-top: 32px;
    margin-left: 32px;
    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    gap: 32px;
  }
</style>