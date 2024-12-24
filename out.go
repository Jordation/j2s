package j2s

type Manufacturer struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Link string `json:"link"`
}
type PriceRange struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}
type Hardpoints struct {
	SubType   string      `json:"sub_type"`
	Position  any         `json:"position"`
	MaxSize   float64     `json:"max_size"`
	ClassName string      `json:"class_name"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	MinSize   float64     `json:"min_size"`
	Item      Item        `json:"item"`
	Children  []*Children `json:"children"`
	Health    any         `json:"health"`
}
type Thruster struct {
	Type                      string  `json:"type"`
	ThrustCapacity            float64 `json:"thrust_capacity"`
	MinHealthThrustMultiplier float64 `json:"min_health_thrust_multiplier"`
	FuelBurnPer10kNewton      float64 `json:"fuel_burn_per_10k_newton"`
}
type PowerPlant struct {
	PowerOutput float64 `json:"power_output"`
}
type Items struct {
	PriceRange       PriceRange `json:"price_range"`
	MaxDiscount      float64    `json:"max_discount"`
	Inventory        float64    `json:"inventory"`
	Buyable          bool       `json:"buyable"`
	Type             string     `json:"type"`
	MaxPremium       float64    `json:"max_premium"`
	Uuid             string     `json:"uuid"`
	BasePrice        float64    `json:"base_price"`
	RefreshRate      float64    `json:"refresh_rate"`
	AutoRestock      bool       `json:"auto_restock"`
	AutoConsume      bool       `json:"auto_consume"`
	Name             string     `json:"name"`
	MaxInventory     float64    `json:"max_inventory"`
	SubType          string     `json:"sub_type"`
	OptimalInventory float64    `json:"optimal_inventory"`
	Version          string     `json:"version"`
	Rentable         bool       `json:"rentable"`
	PriceCalculated  float64    `json:"price_calculated"`
	BasePriceOffset  float64    `json:"base_price_offset"`
	Sellable         bool       `json:"sellable"`
}
type Ship struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}
type Usage struct {
	Vtol        float64 `json:"vtol"`
	Main        float64 `json:"main"`
	Maneuvering float64 `json:"maneuvering"`
	Retro       float64 `json:"retro"`
}
type Shops struct {
	Version      string   `json:"version"`
	Items        []*Items `json:"items"`
	Uuid         string   `json:"uuid"`
	NameRaw      string   `json:"name_raw"`
	Name         string   `json:"name"`
	Position     string   `json:"position"`
	ProfitMargin float64  `json:"profit_margin"`
	Link         string   `json:"link"`
}
type FlightController struct {
	MaxSpeed float64 `json:"max_speed"`
	Pitch    float64 `json:"pitch"`
	Yaw      float64 `json:"yaw"`
	Roll     float64 `json:"roll"`
	ScmSpeed float64 `json:"scm_speed"`
}
type SelfDestruct struct {
	MinPhysRadius float64 `json:"min_phys_radius"`
	Time          float64 `json:"time"`
	Damage        float64 `json:"damage"`
	Radius        float64 `json:"radius"`
	MinRadius     float64 `json:"min_radius"`
	PhysRadius    float64 `json:"phys_radius"`
}
type Insurance struct {
	ClaimTime    float64 `json:"claim_time"`
	ExpediteTime float64 `json:"expedite_time"`
	ExpediteCost float64 `json:"expedite_cost"`
}
type QuantumDrive struct {
	Modes                  []*Modes          `json:"modes"`
	QuantumFuelRequirement float64           `json:"quantum_fuel_requirement"`
	JumpRange              string            `json:"jump_range"`
	DisconnectRange        float64           `json:"disconnect_range"`
	ThermalEnergyDraw      ThermalEnergyDraw `json:"thermal_energy_draw"`
}
type Fuel struct {
	Usage      Usage   `json:"usage"`
	Capacity   float64 `json:"capacity"`
	IntakeRate float64 `json:"intake_rate"`
}
type Ports struct {
	DisplayName     string             `json:"display_name"`
	Position        string             `json:"position"`
	Sizes           Sizes              `json:"sizes"`
	CompatibleTypes []*CompatibleTypes `json:"compatible_types"`
	Tags            []any              `json:"tags"`
	RequiredTags    []any              `json:"required_tags"`
	EquippedItem    any                `json:"equipped_item"`
	Name            string             `json:"name"`
}
type Regeneration struct {
	RequestedRegenPerSec float64 `json:"requested_regen_per_sec"`
	RequestedAmmoLoad    float64 `json:"requested_ammo_load"`
	Cooldown             float64 `json:"cooldown"`
	CostPerBullet        float64 `json:"cost_per_bullet"`
}
type Item struct {
	Emp            Emp            `json:"emp"`
	SelfDestruct   SelfDestruct   `json:"self_destruct"`
	Link           string         `json:"link"`
	Manufacturer   Manufacturer   `json:"manufacturer"`
	FuelIntake     FuelIntake     `json:"fuel_intake"`
	CounterMeasure CounterMeasure `json:"counter_measure"`
	PowerPlant     PowerPlant     `json:"power_plant"`
	Class          string         `json:"class"`
	Grade          string         `json:"grade"`
	MinSize        float64        `json:"min_size"`
	QuantumDrive   QuantumDrive   `json:"quantum_drive"`
	Type           string         `json:"type"`
	MaxMounts      float64        `json:"max_mounts"`
	Thruster       Thruster       `json:"thruster"`
	Shield         Shield         `json:"shield"`
	SubType        string         `json:"sub_type"`
	Version        string         `json:"version"`
	Ports          []*Ports       `json:"ports"`
	UpdatedAt      string         `json:"updated_at"`
	Uuid           string         `json:"uuid"`
	VehicleWeapon  VehicleWeapon  `json:"vehicle_weapon"`
	ClassName      string         `json:"class_name"`
	Cooler         Cooler         `json:"cooler"`
	MaxSize        float64        `json:"max_size"`
	FuelTank       FuelTank       `json:"fuel_tank"`
	Mass           float64        `json:"mass"`
	Inventory      Inventory      `json:"inventory"`
	MaxMissiles    float64        `json:"max_missiles"`
	Name           string         `json:"name"`
	Size           float64        `json:"size"`
}
type MinDamage struct {
	Biochemical float64 `json:"biochemical"`
	Stun        float64 `json:"stun"`
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
}
type Inventory struct {
	Length       float64 `json:"length"`
	Dimension    float64 `json:"dimension"`
	Scu          float64 `json:"scu"`
	ScuConverted float64 `json:"scu_converted"`
	Unit         string  `json:"unit"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
}
type Emp struct {
	SignalInfrared        float64 `json:"signal_infrared"`
	DamageThermal         float64 `json:"damage_thermal"`
	DamageDistortion      float64 `json:"damage_distortion"`
	DamageBiochemical     float64 `json:"damage_biochemical"`
	DamageStun            float64 `json:"damage_stun"`
	DamageEnergy          float64 `json:"damage_energy"`
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	DamagePhysical        float64 `json:"damage_physical"`
	SignalCrossSection    float64 `json:"signal_cross_section"`
}
type Skus struct {
	ImportedAt string  `json:"imported_at"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Available  float64 `json:"available"`
}
type Acceleration struct {
	RetroG       float64 `json:"retro_g"`
	VtolG        float64 `json:"vtol_g"`
	ManeuveringG float64 `json:"maneuvering_g"`
	Main         float64 `json:"main"`
	Retro        float64 `json:"retro"`
	Vtol         float64 `json:"vtol"`
	Maneuvering  float64 `json:"maneuvering"`
	MainG        float64 `json:"main_g"`
}
type Children struct {
	MaxSize     any     `json:"max_size"`
	SubType     string  `json:"sub_type"`
	Children    []any   `json:"children"`
	DisplayName string  `json:"display_name"`
	Position    any     `json:"position"`
	MinSize     any     `json:"min_size"`
	Type        string  `json:"type"`
	Item        Item    `json:"item"`
	ClassName   string  `json:"class_name"`
	Health      float64 `json:"health"`
	DamageMax   float64 `json:"damage_max"`
	Name        string  `json:"name"`
}
type Emission struct {
	EmIdle float64 `json:"em_idle"`
	EmMax  float64 `json:"em_max"`
	Ir     float64 `json:"ir"`
}
type CompatibleTypes struct {
	Type     string   `json:"type"`
	SubTypes []string `json:"sub_types"`
}
type Armor struct {
	SignalCrossSection    float64 `json:"signal_cross_section"`
	DamageStun            float64 `json:"damage_stun"`
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	DamageEnergy          float64 `json:"damage_energy"`
	SignalInfrared        float64 `json:"signal_infrared"`
	DamageDistortion      float64 `json:"damage_distortion"`
	DamageThermal         float64 `json:"damage_thermal"`
	DamageBiochemical     float64 `json:"damage_biochemical"`
	DamagePhysical        float64 `json:"damage_physical"`
}
type Data struct {
	Fuel              Fuel          `json:"fuel"`
	Foci              []string      `json:"foci"`
	Sizes             Sizes         `json:"sizes"`
	PledgeUrl         string        `json:"pledge_url"`
	CargoCapacity     float64       `json:"cargo_capacity"`
	Hardpoints        []*Hardpoints `json:"hardpoints"`
	Health            float64       `json:"health"`
	Parts             []*Parts      `json:"parts"`
	ProductionNote    string        `json:"production_note"`
	Emission          Emission      `json:"emission"`
	Armor             Armor         `json:"armor"`
	Crew              Crew          `json:"crew"`
	Quantum           Quantum       `json:"quantum"`
	Uuid              string        `json:"uuid"`
	Mass              float64       `json:"mass"`
	UpdatedAt         string        `json:"updated_at"`
	Shops             []*Shops      `json:"shops"`
	Description       string        `json:"description"`
	Name              string        `json:"name"`
	SizeClass         float64       `json:"size_class"`
	ProductionStatus  string        `json:"production_status"`
	Slug              string        `json:"slug"`
	Type              string        `json:"type"`
	Speed             Speed         `json:"speed"`
	Size              string        `json:"size"`
	Components        []*Components `json:"components"`
	ClassName         string        `json:"class_name"`
	Msrp              float64       `json:"msrp"`
	ShieldHp          float64       `json:"shield_hp"`
	Version           string        `json:"version"`
	Skus              []*Skus       `json:"skus"`
	PersonalInventory float64       `json:"personal_inventory"`
	Id                float64       `json:"id"`
	Manufacturer      Manufacturer  `json:"manufacturer"`
	Insurance         Insurance     `json:"insurance"`
	ChassisId         float64       `json:"chassis_id"`
	VehicleInventory  float64       `json:"vehicle_inventory"`
	Loaner            []any         `json:"loaner"`
	Agility           Agility       `json:"agility"`
}
type Speed struct {
	MaxToZero float64 `json:"max_to_zero"`
	Scm       float64 `json:"scm"`
	Max       float64 `json:"max"`
	ZeroToScm float64 `json:"zero_to_scm"`
	ZeroToMax float64 `json:"zero_to_max"`
	ScmToZero float64 `json:"scm_to_zero"`
}
type Shield struct {
	ReallocationRate float64    `json:"reallocation_rate"`
	MaxShieldHealth  float64    `json:"max_shield_health"`
	MaxShieldRegen   float64    `json:"max_shield_regen"`
	DecayRatio       float64    `json:"decay_ratio"`
	RegenDelay       RegenDelay `json:"regen_delay"`
	MaxReallocation  float64    `json:"max_reallocation"`
}
type Agility struct {
	Roll         float64      `json:"roll"`
	Acceleration Acceleration `json:"acceleration"`
	Pitch        float64      `json:"pitch"`
	Yaw          float64      `json:"yaw"`
}
type DamageFalloffs struct {
	MinDistance MinDistance `json:"min_distance"`
	PerMeter    PerMeter    `json:"per_meter"`
	MinDamage   MinDamage   `json:"min_damage"`
}
type Ammunition struct {
	Uuid           string         `json:"uuid"`
	Size           float64        `json:"size"`
	Lifetime       float64        `json:"lifetime"`
	Speed          float64        `json:"speed"`
	Range          float64        `json:"range"`
	Piercability   Piercability   `json:"piercability"`
	DamageFalloffs DamageFalloffs `json:"damage_falloffs"`
}
type Missile struct {
	LockRangeMin       float64    `json:"lock_range_min"`
	FuelTankSize       float64    `json:"fuel_tank_size"`
	ExplosionRadiusMax float64    `json:"explosion_radius_max"`
	LockAngle          float64    `json:"lock_angle"`
	ExplosionRadiusMin float64    `json:"explosion_radius_min"`
	DamageTotal        float64    `json:"damage_total"`
	LockTime           float64    `json:"lock_time"`
	TrackingSignalMin  float64    `json:"tracking_signal_min"`
	Speed              float64    `json:"speed"`
	Damages            []*Damages `json:"damages"`
	ClusterSize        float64    `json:"cluster_size"`
	SignalType         string     `json:"signal_type"`
	LockRangeMax       float64    `json:"lock_range_max"`
}
type PerMeter struct {
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
	Biochemical float64 `json:"biochemical"`
	Stun        float64 `json:"stun"`
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
}
type FuelIntake struct {
	FuelPushRate float64 `json:"fuel_push_rate"`
	MinimumRate  float64 `json:"minimum_rate"`
}
type Quantum struct {
	QuantumSpeed        float64 `json:"quantum_speed"`
	QuantumSpoolTime    float64 `json:"quantum_spool_time"`
	QuantumFuelCapacity float64 `json:"quantum_fuel_capacity"`
	QuantumRange        float64 `json:"quantum_range"`
}
type MinDistance struct {
	Biochemical float64 `json:"biochemical"`
	Stun        float64 `json:"stun"`
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
}
type RegenDelay struct {
	Downed float64 `json:"downed"`
	Damage float64 `json:"damage"`
}
type Parts struct {
	DisplayName string      `json:"display_name"`
	DamageMax   float64     `json:"damage_max"`
	Children    []*Children `json:"children"`
	Name        string      `json:"name"`
}
type CounterMeasure struct {
	Range         float64    `json:"range"`
	DamagePerShot float64    `json:"damage_per_shot"`
	Modes         []*Modes   `json:"modes"`
	Damages       []any      `json:"damages"`
	Ammunition    Ammunition `json:"ammunition"`
	Regeneration  any        `json:"regeneration"`
	Type          any        `json:"type"`
	Capacity      float64    `json:"capacity"`
	Class         any        `json:"class"`
}
type Components struct {
	Details        string  `json:"details"`
	Manufacturer   string  `json:"manufacturer"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Size           string  `json:"size"`
	Quantity       float64 `json:"quantity"`
	ComponentClass string  `json:"component_class"`
	Type           string  `json:"type"`
	Mounts         float64 `json:"mounts"`
	ComponentSize  string  `json:"component_size"`
}
type Piercability struct {
	DamageFalloffLevel2     float64 `json:"damage_falloff_level_2"`
	DamageFalloffLevel3     float64 `json:"damage_falloff_level_3"`
	MaxPenetrationThickness float64 `json:"max_penetration_thickness"`
	DamageFalloffLevel1     float64 `json:"damage_falloff_level_1"`
}
type FuelTank struct {
	FillRate  float64 `json:"fill_rate"`
	DrainRate float64 `json:"drain_rate"`
	Capacity  float64 `json:"capacity"`
}
type VehicleWeapon struct {
	Class         any          `json:"class"`
	Capacity      float64      `json:"capacity"`
	Damages       []*Damages   `json:"damages"`
	Ammunition    Ammunition   `json:"ammunition"`
	Range         float64      `json:"range"`
	Modes         []*Modes     `json:"modes"`
	Type          string       `json:"type"`
	DamagePerShot float64      `json:"damage_per_shot"`
	Regeneration  Regeneration `json:"regeneration"`
}
type Sizes struct {
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	Beam   float64 `json:"beam"`
}
type ThermalEnergyDraw struct {
	InFlight     float64 `json:"in_flight"`
	RampDown     float64 `json:"ramp_down"`
	PostRampDown float64 `json:"post_ramp_down"`
	PreRampUp    float64 `json:"pre_ramp_up"`
	RampUp       float64 `json:"ramp_up"`
}
type Modes struct {
	CalibrationRate              float64 `json:"calibration_rate"`
	PelletsPerShot               float64 `json:"pellets_per_shot"`
	DamagePerSecond              float64 `json:"damage_per_second"`
	StageOneAccelRate            float64 `json:"stage_one_accel_rate"`
	EngageSpeed                  float64 `json:"engage_speed"`
	InterdictionEffectTime       float64 `json:"interdiction_effect_time"`
	CalibrationWarningAngleLimit float64 `json:"calibration_warning_angle_limit"`
	DriveSpeed                   float64 `json:"drive_speed"`
	Mode                         string  `json:"mode"`
	AmmoPerShot                  float64 `json:"ammo_per_shot"`
	Type                         string  `json:"type"`
	StageTwoAccelRate            float64 `json:"stage_two_accel_rate"`
	MinCalibrationRequirement    float64 `json:"min_calibration_requirement"`
	SpoolUpTime                  float64 `json:"spool_up_time"`
	CalibrationProcessAngleLimit float64 `json:"calibration_process_angle_limit"`
	Rpm                          float64 `json:"rpm"`
	MaxCalibrationRequirement    float64 `json:"max_calibration_requirement"`
	CooldownTime                 float64 `json:"cooldown_time"`
}
type Cooler struct {
	CoolingRate           float64 `json:"cooling_rate"`
	SuppressionIrFactor   float64 `json:"suppression_ir_factor"`
	SuppressionHeatFactor float64 `json:"suppression_heat_factor"`
}
type Damages struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Damage float64 `json:"damage"`
}
type Meta struct {
	ProcessedAt    string   `json:"processed_at"`
	ValidRelations []string `json:"valid_relations"`
}
type Crew struct {
	Min       float64 `json:"min"`
	Max       any     `json:"max"`
	Weapon    float64 `json:"weapon"`
	Operation any     `json:"operation"`
}
