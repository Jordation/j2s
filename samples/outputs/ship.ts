interface summary {
  meta: meta;
  data: data;
}

interface acceleration {
  maneuveringG: number;
  main: number;
  retro: number;
  vtol: number;
  maneuvering: number;
  mainG: number;
  retroG: number;
  vtolG: number;
}

interface agility {
  roll: number;
  acceleration: acceleration;
  pitch: number;
  yaw: number;
}

interface ammunition {
  damageFalloffs: damageFalloffs;
  uuid: string;
  size: number;
  lifetime: number;
  speed: number;
  range: number;
  piercability: piercability;
}

interface armor {
  signalInfrared: number;
  damageDistortion: number;
  damageThermal: number;
  damageBiochemical: number;
  damageEnergy: number;
  damageStun: number;
  signalElectromagnetic: number;
  damagePhysical: number;
  signalCrossSection: number;
}

interface children {
  type: string;
  minSize: null;
  name: string;
  children: any[];
  health: number;
  position: null;
  maxSize: null;
  subType: string;
  displayName: string;
  damageMax: number;
  item: item;
  className: string;
}

interface compatibleTypes {
  type: string;
  subTypes: string[];
}

interface components {
  category: string;
  manufacturer: string;
  componentSize: string;
  size: string;
  mounts: number;
  details: string;
  quantity: number;
  componentClass: string;
  type: string;
  name: string;
}

interface cooler {
  coolingRate: number;
  suppressionIrFactor: number;
  suppressionHeatFactor: number;
}

interface counterMeasure {
  modes: object[];
  class: null;
  range: number;
  regeneration: null;
  ammunition: ammunition;
  damages: any[];
  type: null;
  capacity: number;
  damagePerShot: number;
}

interface crew {
  min: number;
  max: null;
  weapon: number;
  operation: null;
}

interface damageFalloffs {
  minDistance: minDistance;
  perMeter: perMeter;
  minDamage: minDamage;
}

interface damages {
  type: string;
  name: string;
  damage: number;
}

interface data {
  foci: string[];
  emission: emission;
  personalInventory: number;
  skus: object[];
  parts: object[];
  components: object[];
  pledgeUrl: string;
  hardpoints: object[];
  vehicleInventory: number;
  updatedAt: string;
  crew: crew;
  sizeClass: number;
  insurance: insurance;
  size: string;
  mass: number;
  productionStatus: string;
  chassisId: number;
  msrp: number;
  cargoCapacity: number;
  version: string;
  manufacturer: manufacturer;
  shops: object[];
  name: string;
  speed: speed;
  armor: armor;
  loaner: any[];
  sizes: sizes;
  uuid: string;
  health: number;
  type: string;
  productionNote: string;
  agility: agility;
  fuel: fuel;
  description: string;
  slug: string;
  quantum: quantum;
  id: number;
  shieldHp: number;
  className: string;
}

interface emission {
  emIdle: number;
  emMax: number;
  ir: number;
}

interface emp {
  damageBiochemical: number;
  signalElectromagnetic: number;
  damageDistortion: number;
  damageThermal: number;
  signalInfrared: number;
  signalCrossSection: number;
  damagePhysical: number;
  damageEnergy: number;
  damageStun: number;
}

interface flightController {
  scmSpeed: number;
  maxSpeed: number;
  pitch: number;
  yaw: number;
  roll: number;
}

interface fuel {
  usage: usage;
  capacity: number;
  intakeRate: number;
}

interface fuelIntake {
  minimumRate: number;
  fuelPushRate: number;
}

interface fuelTank {
  fillRate: number;
  drainRate: number;
  capacity: number;
}

interface hardpoints {
  health: null;
  name: string;
  maxSize: number;
  type: string;
  subType: string;
  position: null;
  minSize: number;
  item: item;
  children: object[];
  className: string;
}

interface insurance {
  claimTime: number;
  expediteTime: number;
  expediteCost: number;
}

interface inventory {
  length: number;
  dimension: number;
  scu: number;
  scuConverted: number;
  unit: string;
  width: number;
  height: number;
}

interface item {
  flightController: flightController;
  inventory: inventory;
  grade: string;
  mass: number;
  class: string;
  counterMeasure: counterMeasure;
  minSize: number;
  ports: object[];
  fuelIntake: fuelIntake;
  updatedAt: string;
  name: string;
  manufacturer: manufacturer;
  maxMounts: number;
  maxMissiles: number;
  selfDestruct: selfDestruct;
  className: string;
  subType: string;
  emp: emp;
  quantumDrive: quantumDrive;
  uuid: string;
  version: string;
  link: string;
  size: number;
  type: string;
  maxSize: number;
  cooler: cooler;
  powerPlant: powerPlant;
  shield: shield;
  fuelTank: fuelTank;
  thruster: thruster;
}

interface items {
  refreshRate: number;
  basePrice: number;
  autoRestock: boolean;
  basePriceOffset: number;
  maxDiscount: number;
  optimalInventory: number;
  type: string;
  version: string;
  subType: string;
  priceRange: priceRange;
  maxInventory: number;
  name: string;
  buyable: boolean;
  sellable: boolean;
  priceCalculated: number;
  autoConsume: boolean;
  rentable: boolean;
  maxPremium: number;
  uuid: string;
  inventory: number;
}

interface manufacturer {
  name: string;
  code: string;
  link: string;
}

interface meta {
  validRelations: string[];
  processedAt: string;
}

interface minDamage {
  stun: number;
  physical: number;
  energy: number;
  distortion: number;
  thermal: number;
  biochemical: number;
}

interface minDistance {
  thermal: number;
  biochemical: number;
  stun: number;
  physical: number;
  energy: number;
  distortion: number;
}

interface missile {
  lockTime: number;
  fuelTankSize: number;
  explosionRadiusMax: number;
  explosionRadiusMin: number;
  damageTotal: number;
  lockRangeMin: number;
  trackingSignalMin: number;
  clusterSize: number;
  signalType: string;
  lockRangeMax: number;
  lockAngle: number;
  speed: number;
  damages: object[];
}

interface modes {
  stageOneAccelRate: number;
  calibrationRate: number;
  maxCalibrationRequirement: number;
  stageTwoAccelRate: number;
  calibrationWarningAngleLimit: number;
  type: string;
  damagePerSecond: number;
  mode: string;
  cooldownTime: number;
  pelletsPerShot: number;
  interdictionEffectTime: number;
  calibrationProcessAngleLimit: number;
  minCalibrationRequirement: number;
  spoolUpTime: number;
  driveSpeed: number;
  rpm: number;
  ammoPerShot: number;
  engageSpeed: number;
}

interface parts {
  name: string;
  displayName: string;
  damageMax: number;
  children: object[];
}

interface perMeter {
  thermal: number;
  biochemical: number;
  stun: number;
  physical: number;
  energy: number;
  distortion: number;
}

interface piercability {
  damageFalloffLevel2: number;
  damageFalloffLevel3: number;
  maxPenetrationThickness: number;
  damageFalloffLevel1: number;
}

interface ports {
  position: string;
  tags: any[];
  requiredTags: any[];
  equippedItem: null;
  name: string;
  displayName: string;
  sizes: sizes;
  compatibleTypes: object[];
}

interface powerPlant {
  powerOutput: number;
}

interface priceRange {
  min: number;
  max: number;
}

interface quantum {
  quantumRange: number;
  quantumSpeed: number;
  quantumSpoolTime: number;
  quantumFuelCapacity: number;
}

interface quantumDrive {
  disconnectRange: number;
  thermalEnergyDraw: thermalEnergyDraw;
  modes: object[];
  quantumFuelRequirement: number;
  jumpRange: string;
}

interface regenDelay {
  downed: number;
  damage: number;
}

interface regeneration {
  requestedRegenPerSec: number;
  requestedAmmoLoad: number;
  cooldown: number;
  costPerBullet: number;
}

interface selfDestruct {
  radius: number;
  minRadius: number;
  physRadius: number;
  minPhysRadius: number;
  time: number;
  damage: number;
}

interface shield {
  regenDelay: regenDelay;
  maxReallocation: number;
  reallocationRate: number;
  maxShieldHealth: number;
  maxShieldRegen: number;
  decayRatio: number;
}

interface shops {
  link: string;
  version: string;
  items: object[];
  uuid: string;
  nameRaw: string;
  name: string;
  position: string;
  profitMargin: number;
}

interface sizes {
  beam: number;
  height: number;
  min: number;
  max: number;
  length: number;
}

interface skus {
  title: string;
  price: number;
  available: number;
  importedAt: string;
}

interface speed {
  scm: number;
  max: number;
  zeroToScm: number;
  zeroToMax: number;
  scmToZero: number;
  maxToZero: number;
}

interface thermalEnergyDraw {
  inFlight: number;
  rampDown: number;
  postRampDown: number;
  preRampUp: number;
  rampUp: number;
}

interface thruster {
  thrustCapacity: number;
  minHealthThrustMultiplier: number;
  fuelBurnPer10kNewton: number;
  type: string;
}

interface usage {
  vtol: number;
  main: number;
  maneuvering: number;
  retro: number;
}

interface vehicleWeapon {
  capacity: number;
  ammunition: ammunition;
  modes: object[];
  range: number;
  damagePerShot: number;
  damages: object[];
  regeneration: regeneration;
  class: null;
  type: string;
}

