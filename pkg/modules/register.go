package modules

import (
	"fmt"
	"sync"

	"github.com/DIMO-Network/model-garage/pkg/autopi"
	"github.com/DIMO-Network/model-garage/pkg/compass"
	"github.com/DIMO-Network/model-garage/pkg/defaultmodule"
	"github.com/DIMO-Network/model-garage/pkg/hashdog"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/tesla"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// Register modules to the global registries
	RegisterDefaultModules(SignalRegistry, CloudEventRegistry, FingerprintRegistry)
}

// Source addresses and global registries for different modules.
var (
	// AutoPiSource is the Ethereum address for the AutoPi connection.
	AutoPiSource = common.HexToAddress("0x5e31bBc786D7bEd95216383787deA1ab0f1c1897")

	// RuptelaSource is the Ethereum address for the Ruptela connection.
	RuptelaSource = common.HexToAddress("0xF26421509Efe92861a587482100c6d728aBf1CD0")

	// HashDogSource is the Ethereum address for the HashDog connection.
	HashDogSource = common.HexToAddress("0x4c674ddE8189aEF6e3b58F5a36d7438b2b1f6Bc2")

	// CompassSource is the Ethereum address for the Compass IOT connection.
	CompassSource = common.HexToAddress("0x55BF1c27d468314Ea119CF74979E2b59F962295c")

	// TeslaSource is the Ethereum address for the Tesla connection.
	TeslaSource = common.HexToAddress("0xc4035Fecb1cc906130423EF05f9C20977F643722")

	// SignalRegistry stores signal modules.
	SignalRegistry = NewModuleRegistry[SignalModule]()

	// CloudEventRegistry stores cloud event modules.
	CloudEventRegistry = NewModuleRegistry[CloudEventModule]()

	// FingerprintRegistry stores fingerprint modules.
	FingerprintRegistry = NewModuleRegistry[FingerprintModule]()
)

// RegisterDefaultModules registers all the default module implementations
// into the provided registries.
func RegisterDefaultModules(
	signalReg *ModuleRegistry[SignalModule],
	cloudEventReg *ModuleRegistry[CloudEventModule],
	fingerprintReg *ModuleRegistry[FingerprintModule],
) {
	// AutoPi
	autoPiModule := &autopi.Module{}
	signalReg.Override(AutoPiSource.String(), autoPiModule)
	cloudEventReg.Override(AutoPiSource.String(), autoPiModule)
	fingerprintReg.Override(AutoPiSource.String(), autoPiModule)

	// Ruptela
	ruptelaModule := &ruptela.Module{}
	signalReg.Override(RuptelaSource.String(), ruptelaModule)
	cloudEventReg.Override(RuptelaSource.String(), ruptelaModule)
	fingerprintReg.Override(RuptelaSource.String(), ruptelaModule)

	// HashDog
	hashDogModule := &hashdog.Module{}
	signalReg.Override(HashDogSource.String(), hashDogModule)
	cloudEventReg.Override(HashDogSource.String(), hashDogModule)
	fingerprintReg.Override(HashDogSource.String(), hashDogModule)

	// Compass IOT
	compassModule := &compass.Module{}
	signalReg.Override(CompassSource.String(), compassModule)
	cloudEventReg.Override(CompassSource.String(), compassModule)
	fingerprintReg.Override(CompassSource.String(), compassModule)

	// Tesla
	teslaModule := &tesla.Module{}
	signalReg.Override(TeslaSource.String(), teslaModule)
	cloudEventReg.Override(TeslaSource.String(), teslaModule)
	fingerprintReg.Override(TeslaSource.String(), teslaModule)

	// Default module (empty source)
	defaultModule := &defaultmodule.Module{}
	signalReg.Override("", defaultModule)
	cloudEventReg.Override("", defaultModule)
	fingerprintReg.Override("", defaultModule)
}

// ModuleRegistry is a generic registry for storing and retrieving modules.
type ModuleRegistry[T any] struct {
	mu      sync.RWMutex
	modules map[string]T
}

// NewModuleRegistry creates a new module registry.
func NewModuleRegistry[T any]() *ModuleRegistry[T] {
	return &ModuleRegistry[T]{
		modules: make(map[string]T),
	}
}

// Register adds a module to the registry.
// Returns an error if a module with the same source is already registered.
func (r *ModuleRegistry[T]) Register(source string, module T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.modules[source]; ok {
		return fmt.Errorf("module '%s' already registered", source)
	}
	r.modules[source] = module
	return nil
}

// Override registers or replaces a module in the registry.
func (r *ModuleRegistry[T]) Override(source string, module T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.modules[source] = module
}

// Get retrieves a module from the registry.
func (r *ModuleRegistry[T]) Get(source string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	module, ok := r.modules[source]
	return module, ok
}

// GetSources returns all registered sources.
func (r *ModuleRegistry[T]) GetSources() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sources := make([]string, 0, len(r.modules))
	for source := range r.modules {
		sources = append(sources, source)
	}
	return sources
}
