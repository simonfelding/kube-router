// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package proxy

import (
	"github.com/moby/ipvs"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"net"
	"sync"
)

// Ensure, that LinuxNetworkingMock does implement LinuxNetworking.
// If this is not the case, regenerate this file with moq.
var _ LinuxNetworking = &LinuxNetworkingMock{}

// LinuxNetworkingMock is a mock implementation of LinuxNetworking.
//
//	func TestSomethingThatUsesLinuxNetworking(t *testing.T) {
//
//		// make and configure a mocked LinuxNetworking
//		mockedLinuxNetworking := &LinuxNetworkingMock{
//			configureContainerForDSRFunc: func(vip string, endpointIP string, containerID string, pid int, hostNetworkNamespaceHandle netns.NsHandle) error {
//				panic("mock out the configureContainerForDSR method")
//			},
//			findIfaceLinkForPidFunc: func(pid int) (int, error) {
//				panic("mock out the findIfaceLinkForPid method")
//			},
//			getContainerPidWithCRIFunc: func(runtimeEndpoint string, containerID string) (int, error) {
//				panic("mock out the getContainerPidWithCRI method")
//			},
//			getContainerPidWithDockerFunc: func(containerID string) (int, error) {
//				panic("mock out the getContainerPidWithDocker method")
//			},
//			getKubeDummyInterfaceFunc: func() (netlink.Link, error) {
//				panic("mock out the getKubeDummyInterface method")
//			},
//			ipAddrAddFunc: func(iface netlink.Link, ip string, nodeIP string, addRoute bool) error {
//				panic("mock out the ipAddrAdd method")
//			},
//			ipAddrDelFunc: func(iface netlink.Link, ip string, nodeIP string) error {
//				panic("mock out the ipAddrDel method")
//			},
//			ipvsAddFWMarkServiceFunc: func(svcs []*ipvs.Service, fwMark uint32, family uint16, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) (*ipvs.Service, error) {
//				panic("mock out the ipvsAddFWMarkService method")
//			},
//			ipvsAddServerFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
//				panic("mock out the ipvsAddServer method")
//			},
//			ipvsAddServiceFunc: func(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) ([]*ipvs.Service, *ipvs.Service, error) {
//				panic("mock out the ipvsAddService method")
//			},
//			ipvsDelDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
//				panic("mock out the ipvsDelDestination method")
//			},
//			ipvsDelServiceFunc: func(ipvsSvc *ipvs.Service) error {
//				panic("mock out the ipvsDelService method")
//			},
//			ipvsGetDestinationsFunc: func(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error) {
//				panic("mock out the ipvsGetDestinations method")
//			},
//			ipvsGetServicesFunc: func() ([]*ipvs.Service, error) {
//				panic("mock out the ipvsGetServices method")
//			},
//			ipvsNewDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
//				panic("mock out the ipvsNewDestination method")
//			},
//			ipvsNewServiceFunc: func(ipvsSvc *ipvs.Service) error {
//				panic("mock out the ipvsNewService method")
//			},
//			ipvsUpdateDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
//				panic("mock out the ipvsUpdateDestination method")
//			},
//			ipvsUpdateServiceFunc: func(ipvsSvc *ipvs.Service) error {
//				panic("mock out the ipvsUpdateService method")
//			},
//			setupPolicyRoutingForDSRFunc: func(setupIPv4 bool, setupIPv6 bool) error {
//				panic("mock out the setupPolicyRoutingForDSR method")
//			},
//			setupRoutesForExternalIPForDSRFunc: func(serviceInfo serviceInfoMap, setupIPv4 bool, setupIPv6 bool) error {
//				panic("mock out the setupRoutesForExternalIPForDSR method")
//			},
//		}
//
//		// use mockedLinuxNetworking in code that requires LinuxNetworking
//		// and then make assertions.
//
//	}
type LinuxNetworkingMock struct {
	// configureContainerForDSRFunc mocks the configureContainerForDSR method.
	configureContainerForDSRFunc func(vip string, endpointIP string, containerID string, pid int, hostNetworkNamespaceHandle netns.NsHandle) error

	// findIfaceLinkForPidFunc mocks the findIfaceLinkForPid method.
	findIfaceLinkForPidFunc func(pid int) (int, error)

	// getContainerPidWithCRIFunc mocks the getContainerPidWithCRI method.
	getContainerPidWithCRIFunc func(runtimeEndpoint string, containerID string) (int, error)

	// getContainerPidWithDockerFunc mocks the getContainerPidWithDocker method.
	getContainerPidWithDockerFunc func(containerID string) (int, error)

	// getKubeDummyInterfaceFunc mocks the getKubeDummyInterface method.
	getKubeDummyInterfaceFunc func() (netlink.Link, error)

	// ipAddrAddFunc mocks the ipAddrAdd method.
	ipAddrAddFunc func(iface netlink.Link, ip string, nodeIP string, addRoute bool) error

	// ipAddrDelFunc mocks the ipAddrDel method.
	ipAddrDelFunc func(iface netlink.Link, ip string, nodeIP string) error

	// ipvsAddFWMarkServiceFunc mocks the ipvsAddFWMarkService method.
	ipvsAddFWMarkServiceFunc func(svcs []*ipvs.Service, fwMark uint32, family uint16, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) (*ipvs.Service, error)

	// ipvsAddServerFunc mocks the ipvsAddServer method.
	ipvsAddServerFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsAddServiceFunc mocks the ipvsAddService method.
	ipvsAddServiceFunc func(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) ([]*ipvs.Service, *ipvs.Service, error)

	// ipvsDelDestinationFunc mocks the ipvsDelDestination method.
	ipvsDelDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsDelServiceFunc mocks the ipvsDelService method.
	ipvsDelServiceFunc func(ipvsSvc *ipvs.Service) error

	// ipvsGetDestinationsFunc mocks the ipvsGetDestinations method.
	ipvsGetDestinationsFunc func(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error)

	// ipvsGetServicesFunc mocks the ipvsGetServices method.
	ipvsGetServicesFunc func() ([]*ipvs.Service, error)

	// ipvsNewDestinationFunc mocks the ipvsNewDestination method.
	ipvsNewDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsNewServiceFunc mocks the ipvsNewService method.
	ipvsNewServiceFunc func(ipvsSvc *ipvs.Service) error

	// ipvsUpdateDestinationFunc mocks the ipvsUpdateDestination method.
	ipvsUpdateDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsUpdateServiceFunc mocks the ipvsUpdateService method.
	ipvsUpdateServiceFunc func(ipvsSvc *ipvs.Service) error

	// setupPolicyRoutingForDSRFunc mocks the setupPolicyRoutingForDSR method.
	setupPolicyRoutingForDSRFunc func(setupIPv4 bool, setupIPv6 bool) error

	// setupRoutesForExternalIPForDSRFunc mocks the setupRoutesForExternalIPForDSR method.
	setupRoutesForExternalIPForDSRFunc func(serviceInfo serviceInfoMap, setupIPv4 bool, setupIPv6 bool) error

	// calls tracks calls to the methods.
	calls struct {
		// configureContainerForDSR holds details about calls to the configureContainerForDSR method.
		configureContainerForDSR []struct {
			// Vip is the vip argument value.
			Vip string
			// EndpointIP is the endpointIP argument value.
			EndpointIP string
			// ContainerID is the containerID argument value.
			ContainerID string
			// Pid is the pid argument value.
			Pid int
			// HostNetworkNamespaceHandle is the hostNetworkNamespaceHandle argument value.
			HostNetworkNamespaceHandle netns.NsHandle
		}
		// findIfaceLinkForPid holds details about calls to the findIfaceLinkForPid method.
		findIfaceLinkForPid []struct {
			// Pid is the pid argument value.
			Pid int
		}
		// getContainerPidWithCRI holds details about calls to the getContainerPidWithCRI method.
		getContainerPidWithCRI []struct {
			// RuntimeEndpoint is the runtimeEndpoint argument value.
			RuntimeEndpoint string
			// ContainerID is the containerID argument value.
			ContainerID string
		}
		// getContainerPidWithDocker holds details about calls to the getContainerPidWithDocker method.
		getContainerPidWithDocker []struct {
			// ContainerID is the containerID argument value.
			ContainerID string
		}
		// getKubeDummyInterface holds details about calls to the getKubeDummyInterface method.
		getKubeDummyInterface []struct {
		}
		// ipAddrAdd holds details about calls to the ipAddrAdd method.
		ipAddrAdd []struct {
			// Iface is the iface argument value.
			Iface netlink.Link
			// IP is the ip argument value.
			IP string
			// NodeIP is the nodeIP argument value.
			NodeIP string
			// AddRoute is the addRoute argument value.
			AddRoute bool
		}
		// ipAddrDel holds details about calls to the ipAddrDel method.
		ipAddrDel []struct {
			// Iface is the iface argument value.
			Iface netlink.Link
			// IP is the ip argument value.
			IP string
			// NodeIP is the nodeIP argument value.
			NodeIP string
		}
		// ipvsAddFWMarkService holds details about calls to the ipvsAddFWMarkService method.
		ipvsAddFWMarkService []struct {
			// Svcs is the svcs argument value.
			Svcs []*ipvs.Service
			// FwMark is the fwMark argument value.
			FwMark uint32
			// Family is the family argument value.
			Family uint16
			// Protocol is the protocol argument value.
			Protocol uint16
			// Port is the port argument value.
			Port uint16
			// Persistent is the persistent argument value.
			Persistent bool
			// PersistentTimeout is the persistentTimeout argument value.
			PersistentTimeout int32
			// Scheduler is the scheduler argument value.
			Scheduler string
			// Flags is the flags argument value.
			Flags schedFlags
		}
		// ipvsAddServer holds details about calls to the ipvsAddServer method.
		ipvsAddServer []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsAddService holds details about calls to the ipvsAddService method.
		ipvsAddService []struct {
			// Svcs is the svcs argument value.
			Svcs []*ipvs.Service
			// Vip is the vip argument value.
			Vip net.IP
			// Protocol is the protocol argument value.
			Protocol uint16
			// Port is the port argument value.
			Port uint16
			// Persistent is the persistent argument value.
			Persistent bool
			// PersistentTimeout is the persistentTimeout argument value.
			PersistentTimeout int32
			// Scheduler is the scheduler argument value.
			Scheduler string
			// Flags is the flags argument value.
			Flags schedFlags
		}
		// ipvsDelDestination holds details about calls to the ipvsDelDestination method.
		ipvsDelDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsDelService holds details about calls to the ipvsDelService method.
		ipvsDelService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsGetDestinations holds details about calls to the ipvsGetDestinations method.
		ipvsGetDestinations []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsGetServices holds details about calls to the ipvsGetServices method.
		ipvsGetServices []struct {
		}
		// ipvsNewDestination holds details about calls to the ipvsNewDestination method.
		ipvsNewDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsNewService holds details about calls to the ipvsNewService method.
		ipvsNewService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsUpdateDestination holds details about calls to the ipvsUpdateDestination method.
		ipvsUpdateDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsUpdateService holds details about calls to the ipvsUpdateService method.
		ipvsUpdateService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// setupPolicyRoutingForDSR holds details about calls to the setupPolicyRoutingForDSR method.
		setupPolicyRoutingForDSR []struct {
			// SetupIPv4 is the setupIPv4 argument value.
			SetupIPv4 bool
			// SetupIPv6 is the setupIPv6 argument value.
			SetupIPv6 bool
		}
		// setupRoutesForExternalIPForDSR holds details about calls to the setupRoutesForExternalIPForDSR method.
		setupRoutesForExternalIPForDSR []struct {
			// ServiceInfo is the serviceInfo argument value.
			ServiceInfo serviceInfoMap
			// SetupIPv4 is the setupIPv4 argument value.
			SetupIPv4 bool
			// SetupIPv6 is the setupIPv6 argument value.
			SetupIPv6 bool
		}
	}
	lockconfigureContainerForDSR       sync.RWMutex
	lockfindIfaceLinkForPid            sync.RWMutex
	lockgetContainerPidWithCRI         sync.RWMutex
	lockgetContainerPidWithDocker      sync.RWMutex
	lockgetKubeDummyInterface          sync.RWMutex
	lockipAddrAdd                      sync.RWMutex
	lockipAddrDel                      sync.RWMutex
	lockipvsAddFWMarkService           sync.RWMutex
	lockipvsAddServer                  sync.RWMutex
	lockipvsAddService                 sync.RWMutex
	lockipvsDelDestination             sync.RWMutex
	lockipvsDelService                 sync.RWMutex
	lockipvsGetDestinations            sync.RWMutex
	lockipvsGetServices                sync.RWMutex
	lockipvsNewDestination             sync.RWMutex
	lockipvsNewService                 sync.RWMutex
	lockipvsUpdateDestination          sync.RWMutex
	lockipvsUpdateService              sync.RWMutex
	locksetupPolicyRoutingForDSR       sync.RWMutex
	locksetupRoutesForExternalIPForDSR sync.RWMutex
}

// configureContainerForDSR calls configureContainerForDSRFunc.
func (mock *LinuxNetworkingMock) configureContainerForDSR(vip string, endpointIP string, containerID string, pid int, hostNetworkNamespaceHandle netns.NsHandle) error {
	if mock.configureContainerForDSRFunc == nil {
		panic("LinuxNetworkingMock.configureContainerForDSRFunc: method is nil but LinuxNetworking.configureContainerForDSR was just called")
	}
	callInfo := struct {
		Vip                        string
		EndpointIP                 string
		ContainerID                string
		Pid                        int
		HostNetworkNamespaceHandle netns.NsHandle
	}{
		Vip:                        vip,
		EndpointIP:                 endpointIP,
		ContainerID:                containerID,
		Pid:                        pid,
		HostNetworkNamespaceHandle: hostNetworkNamespaceHandle,
	}
	mock.lockconfigureContainerForDSR.Lock()
	mock.calls.configureContainerForDSR = append(mock.calls.configureContainerForDSR, callInfo)
	mock.lockconfigureContainerForDSR.Unlock()
	return mock.configureContainerForDSRFunc(vip, endpointIP, containerID, pid, hostNetworkNamespaceHandle)
}

// configureContainerForDSRCalls gets all the calls that were made to configureContainerForDSR.
// Check the length with:
//
//	len(mockedLinuxNetworking.configureContainerForDSRCalls())
func (mock *LinuxNetworkingMock) configureContainerForDSRCalls() []struct {
	Vip                        string
	EndpointIP                 string
	ContainerID                string
	Pid                        int
	HostNetworkNamespaceHandle netns.NsHandle
} {
	var calls []struct {
		Vip                        string
		EndpointIP                 string
		ContainerID                string
		Pid                        int
		HostNetworkNamespaceHandle netns.NsHandle
	}
	mock.lockconfigureContainerForDSR.RLock()
	calls = mock.calls.configureContainerForDSR
	mock.lockconfigureContainerForDSR.RUnlock()
	return calls
}

// findIfaceLinkForPid calls findIfaceLinkForPidFunc.
func (mock *LinuxNetworkingMock) findIfaceLinkForPid(pid int) (int, error) {
	if mock.findIfaceLinkForPidFunc == nil {
		panic("LinuxNetworkingMock.findIfaceLinkForPidFunc: method is nil but LinuxNetworking.findIfaceLinkForPid was just called")
	}
	callInfo := struct {
		Pid int
	}{
		Pid: pid,
	}
	mock.lockfindIfaceLinkForPid.Lock()
	mock.calls.findIfaceLinkForPid = append(mock.calls.findIfaceLinkForPid, callInfo)
	mock.lockfindIfaceLinkForPid.Unlock()
	return mock.findIfaceLinkForPidFunc(pid)
}

// findIfaceLinkForPidCalls gets all the calls that were made to findIfaceLinkForPid.
// Check the length with:
//
//	len(mockedLinuxNetworking.findIfaceLinkForPidCalls())
func (mock *LinuxNetworkingMock) findIfaceLinkForPidCalls() []struct {
	Pid int
} {
	var calls []struct {
		Pid int
	}
	mock.lockfindIfaceLinkForPid.RLock()
	calls = mock.calls.findIfaceLinkForPid
	mock.lockfindIfaceLinkForPid.RUnlock()
	return calls
}

// getContainerPidWithCRI calls getContainerPidWithCRIFunc.
func (mock *LinuxNetworkingMock) getContainerPidWithCRI(runtimeEndpoint string, containerID string) (int, error) {
	if mock.getContainerPidWithCRIFunc == nil {
		panic("LinuxNetworkingMock.getContainerPidWithCRIFunc: method is nil but LinuxNetworking.getContainerPidWithCRI was just called")
	}
	callInfo := struct {
		RuntimeEndpoint string
		ContainerID     string
	}{
		RuntimeEndpoint: runtimeEndpoint,
		ContainerID:     containerID,
	}
	mock.lockgetContainerPidWithCRI.Lock()
	mock.calls.getContainerPidWithCRI = append(mock.calls.getContainerPidWithCRI, callInfo)
	mock.lockgetContainerPidWithCRI.Unlock()
	return mock.getContainerPidWithCRIFunc(runtimeEndpoint, containerID)
}

// getContainerPidWithCRICalls gets all the calls that were made to getContainerPidWithCRI.
// Check the length with:
//
//	len(mockedLinuxNetworking.getContainerPidWithCRICalls())
func (mock *LinuxNetworkingMock) getContainerPidWithCRICalls() []struct {
	RuntimeEndpoint string
	ContainerID     string
} {
	var calls []struct {
		RuntimeEndpoint string
		ContainerID     string
	}
	mock.lockgetContainerPidWithCRI.RLock()
	calls = mock.calls.getContainerPidWithCRI
	mock.lockgetContainerPidWithCRI.RUnlock()
	return calls
}

// getContainerPidWithDocker calls getContainerPidWithDockerFunc.
func (mock *LinuxNetworkingMock) getContainerPidWithDocker(containerID string) (int, error) {
	if mock.getContainerPidWithDockerFunc == nil {
		panic("LinuxNetworkingMock.getContainerPidWithDockerFunc: method is nil but LinuxNetworking.getContainerPidWithDocker was just called")
	}
	callInfo := struct {
		ContainerID string
	}{
		ContainerID: containerID,
	}
	mock.lockgetContainerPidWithDocker.Lock()
	mock.calls.getContainerPidWithDocker = append(mock.calls.getContainerPidWithDocker, callInfo)
	mock.lockgetContainerPidWithDocker.Unlock()
	return mock.getContainerPidWithDockerFunc(containerID)
}

// getContainerPidWithDockerCalls gets all the calls that were made to getContainerPidWithDocker.
// Check the length with:
//
//	len(mockedLinuxNetworking.getContainerPidWithDockerCalls())
func (mock *LinuxNetworkingMock) getContainerPidWithDockerCalls() []struct {
	ContainerID string
} {
	var calls []struct {
		ContainerID string
	}
	mock.lockgetContainerPidWithDocker.RLock()
	calls = mock.calls.getContainerPidWithDocker
	mock.lockgetContainerPidWithDocker.RUnlock()
	return calls
}

// getKubeDummyInterface calls getKubeDummyInterfaceFunc.
func (mock *LinuxNetworkingMock) getKubeDummyInterface() (netlink.Link, error) {
	if mock.getKubeDummyInterfaceFunc == nil {
		panic("LinuxNetworkingMock.getKubeDummyInterfaceFunc: method is nil but LinuxNetworking.getKubeDummyInterface was just called")
	}
	callInfo := struct {
	}{}
	mock.lockgetKubeDummyInterface.Lock()
	mock.calls.getKubeDummyInterface = append(mock.calls.getKubeDummyInterface, callInfo)
	mock.lockgetKubeDummyInterface.Unlock()
	return mock.getKubeDummyInterfaceFunc()
}

// getKubeDummyInterfaceCalls gets all the calls that were made to getKubeDummyInterface.
// Check the length with:
//
//	len(mockedLinuxNetworking.getKubeDummyInterfaceCalls())
func (mock *LinuxNetworkingMock) getKubeDummyInterfaceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockgetKubeDummyInterface.RLock()
	calls = mock.calls.getKubeDummyInterface
	mock.lockgetKubeDummyInterface.RUnlock()
	return calls
}

// ipAddrAdd calls ipAddrAddFunc.
func (mock *LinuxNetworkingMock) ipAddrAdd(iface netlink.Link, ip string, nodeIP string, addRoute bool) error {
	if mock.ipAddrAddFunc == nil {
		panic("LinuxNetworkingMock.ipAddrAddFunc: method is nil but LinuxNetworking.ipAddrAdd was just called")
	}
	callInfo := struct {
		Iface    netlink.Link
		IP       string
		NodeIP   string
		AddRoute bool
	}{
		Iface:    iface,
		IP:       ip,
		NodeIP:   nodeIP,
		AddRoute: addRoute,
	}
	mock.lockipAddrAdd.Lock()
	mock.calls.ipAddrAdd = append(mock.calls.ipAddrAdd, callInfo)
	mock.lockipAddrAdd.Unlock()
	return mock.ipAddrAddFunc(iface, ip, nodeIP, addRoute)
}

// ipAddrAddCalls gets all the calls that were made to ipAddrAdd.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipAddrAddCalls())
func (mock *LinuxNetworkingMock) ipAddrAddCalls() []struct {
	Iface    netlink.Link
	IP       string
	NodeIP   string
	AddRoute bool
} {
	var calls []struct {
		Iface    netlink.Link
		IP       string
		NodeIP   string
		AddRoute bool
	}
	mock.lockipAddrAdd.RLock()
	calls = mock.calls.ipAddrAdd
	mock.lockipAddrAdd.RUnlock()
	return calls
}

// ipAddrDel calls ipAddrDelFunc.
func (mock *LinuxNetworkingMock) ipAddrDel(iface netlink.Link, ip string, nodeIP string) error {
	if mock.ipAddrDelFunc == nil {
		panic("LinuxNetworkingMock.ipAddrDelFunc: method is nil but LinuxNetworking.ipAddrDel was just called")
	}
	callInfo := struct {
		Iface  netlink.Link
		IP     string
		NodeIP string
	}{
		Iface:  iface,
		IP:     ip,
		NodeIP: nodeIP,
	}
	mock.lockipAddrDel.Lock()
	mock.calls.ipAddrDel = append(mock.calls.ipAddrDel, callInfo)
	mock.lockipAddrDel.Unlock()
	return mock.ipAddrDelFunc(iface, ip, nodeIP)
}

// ipAddrDelCalls gets all the calls that were made to ipAddrDel.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipAddrDelCalls())
func (mock *LinuxNetworkingMock) ipAddrDelCalls() []struct {
	Iface  netlink.Link
	IP     string
	NodeIP string
} {
	var calls []struct {
		Iface  netlink.Link
		IP     string
		NodeIP string
	}
	mock.lockipAddrDel.RLock()
	calls = mock.calls.ipAddrDel
	mock.lockipAddrDel.RUnlock()
	return calls
}

// ipvsAddFWMarkService calls ipvsAddFWMarkServiceFunc.
func (mock *LinuxNetworkingMock) ipvsAddFWMarkService(svcs []*ipvs.Service, fwMark uint32, family uint16, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) (*ipvs.Service, error) {
	if mock.ipvsAddFWMarkServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddFWMarkServiceFunc: method is nil but LinuxNetworking.ipvsAddFWMarkService was just called")
	}
	callInfo := struct {
		Svcs              []*ipvs.Service
		FwMark            uint32
		Family            uint16
		Protocol          uint16
		Port              uint16
		Persistent        bool
		PersistentTimeout int32
		Scheduler         string
		Flags             schedFlags
	}{
		Svcs:              svcs,
		FwMark:            fwMark,
		Family:            family,
		Protocol:          protocol,
		Port:              port,
		Persistent:        persistent,
		PersistentTimeout: persistentTimeout,
		Scheduler:         scheduler,
		Flags:             flags,
	}
	mock.lockipvsAddFWMarkService.Lock()
	mock.calls.ipvsAddFWMarkService = append(mock.calls.ipvsAddFWMarkService, callInfo)
	mock.lockipvsAddFWMarkService.Unlock()
	return mock.ipvsAddFWMarkServiceFunc(svcs, fwMark, family, protocol, port, persistent, persistentTimeout, scheduler, flags)
}

// ipvsAddFWMarkServiceCalls gets all the calls that were made to ipvsAddFWMarkService.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsAddFWMarkServiceCalls())
func (mock *LinuxNetworkingMock) ipvsAddFWMarkServiceCalls() []struct {
	Svcs              []*ipvs.Service
	FwMark            uint32
	Family            uint16
	Protocol          uint16
	Port              uint16
	Persistent        bool
	PersistentTimeout int32
	Scheduler         string
	Flags             schedFlags
} {
	var calls []struct {
		Svcs              []*ipvs.Service
		FwMark            uint32
		Family            uint16
		Protocol          uint16
		Port              uint16
		Persistent        bool
		PersistentTimeout int32
		Scheduler         string
		Flags             schedFlags
	}
	mock.lockipvsAddFWMarkService.RLock()
	calls = mock.calls.ipvsAddFWMarkService
	mock.lockipvsAddFWMarkService.RUnlock()
	return calls
}

// ipvsAddServer calls ipvsAddServerFunc.
func (mock *LinuxNetworkingMock) ipvsAddServer(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsAddServerFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddServerFunc: method is nil but LinuxNetworking.ipvsAddServer was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	mock.lockipvsAddServer.Lock()
	mock.calls.ipvsAddServer = append(mock.calls.ipvsAddServer, callInfo)
	mock.lockipvsAddServer.Unlock()
	return mock.ipvsAddServerFunc(ipvsSvc, ipvsDst)
}

// ipvsAddServerCalls gets all the calls that were made to ipvsAddServer.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsAddServerCalls())
func (mock *LinuxNetworkingMock) ipvsAddServerCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	mock.lockipvsAddServer.RLock()
	calls = mock.calls.ipvsAddServer
	mock.lockipvsAddServer.RUnlock()
	return calls
}

// ipvsAddService calls ipvsAddServiceFunc.
func (mock *LinuxNetworkingMock) ipvsAddService(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, persistentTimeout int32, scheduler string, flags schedFlags) ([]*ipvs.Service, *ipvs.Service, error) {
	if mock.ipvsAddServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddServiceFunc: method is nil but LinuxNetworking.ipvsAddService was just called")
	}
	callInfo := struct {
		Svcs              []*ipvs.Service
		Vip               net.IP
		Protocol          uint16
		Port              uint16
		Persistent        bool
		PersistentTimeout int32
		Scheduler         string
		Flags             schedFlags
	}{
		Svcs:              svcs,
		Vip:               vip,
		Protocol:          protocol,
		Port:              port,
		Persistent:        persistent,
		PersistentTimeout: persistentTimeout,
		Scheduler:         scheduler,
		Flags:             flags,
	}
	mock.lockipvsAddService.Lock()
	mock.calls.ipvsAddService = append(mock.calls.ipvsAddService, callInfo)
	mock.lockipvsAddService.Unlock()
	return mock.ipvsAddServiceFunc(svcs, vip, protocol, port, persistent, persistentTimeout, scheduler, flags)
}

// ipvsAddServiceCalls gets all the calls that were made to ipvsAddService.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsAddServiceCalls())
func (mock *LinuxNetworkingMock) ipvsAddServiceCalls() []struct {
	Svcs              []*ipvs.Service
	Vip               net.IP
	Protocol          uint16
	Port              uint16
	Persistent        bool
	PersistentTimeout int32
	Scheduler         string
	Flags             schedFlags
} {
	var calls []struct {
		Svcs              []*ipvs.Service
		Vip               net.IP
		Protocol          uint16
		Port              uint16
		Persistent        bool
		PersistentTimeout int32
		Scheduler         string
		Flags             schedFlags
	}
	mock.lockipvsAddService.RLock()
	calls = mock.calls.ipvsAddService
	mock.lockipvsAddService.RUnlock()
	return calls
}

// ipvsDelDestination calls ipvsDelDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsDelDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsDelDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsDelDestinationFunc: method is nil but LinuxNetworking.ipvsDelDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	mock.lockipvsDelDestination.Lock()
	mock.calls.ipvsDelDestination = append(mock.calls.ipvsDelDestination, callInfo)
	mock.lockipvsDelDestination.Unlock()
	return mock.ipvsDelDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsDelDestinationCalls gets all the calls that were made to ipvsDelDestination.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsDelDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsDelDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	mock.lockipvsDelDestination.RLock()
	calls = mock.calls.ipvsDelDestination
	mock.lockipvsDelDestination.RUnlock()
	return calls
}

// ipvsDelService calls ipvsDelServiceFunc.
func (mock *LinuxNetworkingMock) ipvsDelService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsDelServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsDelServiceFunc: method is nil but LinuxNetworking.ipvsDelService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	mock.lockipvsDelService.Lock()
	mock.calls.ipvsDelService = append(mock.calls.ipvsDelService, callInfo)
	mock.lockipvsDelService.Unlock()
	return mock.ipvsDelServiceFunc(ipvsSvc)
}

// ipvsDelServiceCalls gets all the calls that were made to ipvsDelService.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsDelServiceCalls())
func (mock *LinuxNetworkingMock) ipvsDelServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	mock.lockipvsDelService.RLock()
	calls = mock.calls.ipvsDelService
	mock.lockipvsDelService.RUnlock()
	return calls
}

// ipvsGetDestinations calls ipvsGetDestinationsFunc.
func (mock *LinuxNetworkingMock) ipvsGetDestinations(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error) {
	if mock.ipvsGetDestinationsFunc == nil {
		panic("LinuxNetworkingMock.ipvsGetDestinationsFunc: method is nil but LinuxNetworking.ipvsGetDestinations was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	mock.lockipvsGetDestinations.Lock()
	mock.calls.ipvsGetDestinations = append(mock.calls.ipvsGetDestinations, callInfo)
	mock.lockipvsGetDestinations.Unlock()
	return mock.ipvsGetDestinationsFunc(ipvsSvc)
}

// ipvsGetDestinationsCalls gets all the calls that were made to ipvsGetDestinations.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsGetDestinationsCalls())
func (mock *LinuxNetworkingMock) ipvsGetDestinationsCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	mock.lockipvsGetDestinations.RLock()
	calls = mock.calls.ipvsGetDestinations
	mock.lockipvsGetDestinations.RUnlock()
	return calls
}

// ipvsGetServices calls ipvsGetServicesFunc.
func (mock *LinuxNetworkingMock) ipvsGetServices() ([]*ipvs.Service, error) {
	if mock.ipvsGetServicesFunc == nil {
		panic("LinuxNetworkingMock.ipvsGetServicesFunc: method is nil but LinuxNetworking.ipvsGetServices was just called")
	}
	callInfo := struct {
	}{}
	mock.lockipvsGetServices.Lock()
	mock.calls.ipvsGetServices = append(mock.calls.ipvsGetServices, callInfo)
	mock.lockipvsGetServices.Unlock()
	return mock.ipvsGetServicesFunc()
}

// ipvsGetServicesCalls gets all the calls that were made to ipvsGetServices.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsGetServicesCalls())
func (mock *LinuxNetworkingMock) ipvsGetServicesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockipvsGetServices.RLock()
	calls = mock.calls.ipvsGetServices
	mock.lockipvsGetServices.RUnlock()
	return calls
}

// ipvsNewDestination calls ipvsNewDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsNewDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsNewDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsNewDestinationFunc: method is nil but LinuxNetworking.ipvsNewDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	mock.lockipvsNewDestination.Lock()
	mock.calls.ipvsNewDestination = append(mock.calls.ipvsNewDestination, callInfo)
	mock.lockipvsNewDestination.Unlock()
	return mock.ipvsNewDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsNewDestinationCalls gets all the calls that were made to ipvsNewDestination.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsNewDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsNewDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	mock.lockipvsNewDestination.RLock()
	calls = mock.calls.ipvsNewDestination
	mock.lockipvsNewDestination.RUnlock()
	return calls
}

// ipvsNewService calls ipvsNewServiceFunc.
func (mock *LinuxNetworkingMock) ipvsNewService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsNewServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsNewServiceFunc: method is nil but LinuxNetworking.ipvsNewService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	mock.lockipvsNewService.Lock()
	mock.calls.ipvsNewService = append(mock.calls.ipvsNewService, callInfo)
	mock.lockipvsNewService.Unlock()
	return mock.ipvsNewServiceFunc(ipvsSvc)
}

// ipvsNewServiceCalls gets all the calls that were made to ipvsNewService.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsNewServiceCalls())
func (mock *LinuxNetworkingMock) ipvsNewServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	mock.lockipvsNewService.RLock()
	calls = mock.calls.ipvsNewService
	mock.lockipvsNewService.RUnlock()
	return calls
}

// ipvsUpdateDestination calls ipvsUpdateDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsUpdateDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsUpdateDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsUpdateDestinationFunc: method is nil but LinuxNetworking.ipvsUpdateDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	mock.lockipvsUpdateDestination.Lock()
	mock.calls.ipvsUpdateDestination = append(mock.calls.ipvsUpdateDestination, callInfo)
	mock.lockipvsUpdateDestination.Unlock()
	return mock.ipvsUpdateDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsUpdateDestinationCalls gets all the calls that were made to ipvsUpdateDestination.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsUpdateDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsUpdateDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	mock.lockipvsUpdateDestination.RLock()
	calls = mock.calls.ipvsUpdateDestination
	mock.lockipvsUpdateDestination.RUnlock()
	return calls
}

// ipvsUpdateService calls ipvsUpdateServiceFunc.
func (mock *LinuxNetworkingMock) ipvsUpdateService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsUpdateServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsUpdateServiceFunc: method is nil but LinuxNetworking.ipvsUpdateService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	mock.lockipvsUpdateService.Lock()
	mock.calls.ipvsUpdateService = append(mock.calls.ipvsUpdateService, callInfo)
	mock.lockipvsUpdateService.Unlock()
	return mock.ipvsUpdateServiceFunc(ipvsSvc)
}

// ipvsUpdateServiceCalls gets all the calls that were made to ipvsUpdateService.
// Check the length with:
//
//	len(mockedLinuxNetworking.ipvsUpdateServiceCalls())
func (mock *LinuxNetworkingMock) ipvsUpdateServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	mock.lockipvsUpdateService.RLock()
	calls = mock.calls.ipvsUpdateService
	mock.lockipvsUpdateService.RUnlock()
	return calls
}

// setupPolicyRoutingForDSR calls setupPolicyRoutingForDSRFunc.
func (mock *LinuxNetworkingMock) setupPolicyRoutingForDSR(setupIPv4 bool, setupIPv6 bool) error {
	if mock.setupPolicyRoutingForDSRFunc == nil {
		panic("LinuxNetworkingMock.setupPolicyRoutingForDSRFunc: method is nil but LinuxNetworking.setupPolicyRoutingForDSR was just called")
	}
	callInfo := struct {
		SetupIPv4 bool
		SetupIPv6 bool
	}{
		SetupIPv4: setupIPv4,
		SetupIPv6: setupIPv6,
	}
	mock.locksetupPolicyRoutingForDSR.Lock()
	mock.calls.setupPolicyRoutingForDSR = append(mock.calls.setupPolicyRoutingForDSR, callInfo)
	mock.locksetupPolicyRoutingForDSR.Unlock()
	return mock.setupPolicyRoutingForDSRFunc(setupIPv4, setupIPv6)
}

// setupPolicyRoutingForDSRCalls gets all the calls that were made to setupPolicyRoutingForDSR.
// Check the length with:
//
//	len(mockedLinuxNetworking.setupPolicyRoutingForDSRCalls())
func (mock *LinuxNetworkingMock) setupPolicyRoutingForDSRCalls() []struct {
	SetupIPv4 bool
	SetupIPv6 bool
} {
	var calls []struct {
		SetupIPv4 bool
		SetupIPv6 bool
	}
	mock.locksetupPolicyRoutingForDSR.RLock()
	calls = mock.calls.setupPolicyRoutingForDSR
	mock.locksetupPolicyRoutingForDSR.RUnlock()
	return calls
}

// setupRoutesForExternalIPForDSR calls setupRoutesForExternalIPForDSRFunc.
func (mock *LinuxNetworkingMock) setupRoutesForExternalIPForDSR(serviceInfo serviceInfoMap, setupIPv4 bool, setupIPv6 bool) error {
	if mock.setupRoutesForExternalIPForDSRFunc == nil {
		panic("LinuxNetworkingMock.setupRoutesForExternalIPForDSRFunc: method is nil but LinuxNetworking.setupRoutesForExternalIPForDSR was just called")
	}
	callInfo := struct {
		ServiceInfo serviceInfoMap
		SetupIPv4   bool
		SetupIPv6   bool
	}{
		ServiceInfo: serviceInfo,
		SetupIPv4:   setupIPv4,
		SetupIPv6:   setupIPv6,
	}
	mock.locksetupRoutesForExternalIPForDSR.Lock()
	mock.calls.setupRoutesForExternalIPForDSR = append(mock.calls.setupRoutesForExternalIPForDSR, callInfo)
	mock.locksetupRoutesForExternalIPForDSR.Unlock()
	return mock.setupRoutesForExternalIPForDSRFunc(serviceInfo, setupIPv4, setupIPv6)
}

// setupRoutesForExternalIPForDSRCalls gets all the calls that were made to setupRoutesForExternalIPForDSR.
// Check the length with:
//
//	len(mockedLinuxNetworking.setupRoutesForExternalIPForDSRCalls())
func (mock *LinuxNetworkingMock) setupRoutesForExternalIPForDSRCalls() []struct {
	ServiceInfo serviceInfoMap
	SetupIPv4   bool
	SetupIPv6   bool
} {
	var calls []struct {
		ServiceInfo serviceInfoMap
		SetupIPv4   bool
		SetupIPv6   bool
	}
	mock.locksetupRoutesForExternalIPForDSR.RLock()
	calls = mock.calls.setupRoutesForExternalIPForDSR
	mock.locksetupRoutesForExternalIPForDSR.RUnlock()
	return calls
}
