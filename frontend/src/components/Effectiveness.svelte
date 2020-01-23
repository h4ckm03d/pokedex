<script>
  import EnergyType from "./EnergyType.svelte";
  export let effects;
</script>

<p class="title is-5">
  {#if effects}
    <span>
      <span>
        <EnergyType types={[effects[0].type]} />
      </span>
      {effects[0].value}
    </span>
  {:else}
    <span>N/A</span>
  {/if}
</p>
