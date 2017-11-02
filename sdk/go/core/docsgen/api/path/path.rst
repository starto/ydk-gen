Path
====

.. go:package:: ydk/path
    :synopsis: Path API

.. function:: ExecuteRPC(provider ServiceProvider, entity Entity, filter string, dataTag string, setConfigFlag bool)

    Executes payload converted from entity.

    :param provider: (:go:struct:`ServiceProvider<ydk/types/ServiceProvider>`).
    :param entity: (:go:struct:`Entity<ydk/types/Entity>`).
    :param filter: (``string``) A Go string.
    :param dataTag: (``string``) A Go string.
    :param setConfigFlag: (``bool``) A Go bool.
    :return: A data node representing the result of the executed rpc.
    :rtype: :go:struct:`DataNode<ydk/types/DataNode>`

.. function:: ReadDatanode(filter Entity, readDataNode DataNode)

    Populates entity by reading the top level entity from a given data node

    :param filter: (:go:struct:`Entity<ydk/types/Entity>`)
    :param readDataNode: (:go:struct:`DataNode<ydk/types/DataNode>`)
    :return: The top entity from readDataNode.
    :rtype: :go:struct:`Entity<ydk/types/Entity>`

.. function:: ConnectToNetconfProvider(state *State, repo Repository, address, username, password string, port int)
    
    Connects to NETCONF service provider by creating a connection to the given provider using given address, username, password, and port.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param repo: (:go:struct:`Repository<ydk/types/Repository>`).
    :param address: (``string``) A Go string.
    :param username: (``string``) A Go string.
    :param password: (``string``) A Go string.
    :param port: (``int``) An integer.
    :return: The connected service provider.
    :rtype: :go:struct:`CServiceProvider<ydk/types/CServiceProvider>`

.. function:: DisconnectFromNetconfProvider(provider CServiceProvider)

    Disconnects from NETCONF device and frees the given service provider

    :param: provider: (:go:struct:`CServiceProvider<ydk/cgopath/CServiceProvider>`) A service provider instance.

.. function:: CleanUpErrorState(state *State)
    
    CleanUpErrorState cleans up memory for CState

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution

.. function:: ConnectToRestconfProvider(state *State, path, address, username, password string, port int)
    
    ConnectToRestconfProvider connects to RESTCONF device by creating a connection to the provider using given path, address, username, password, and port.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param path: (``string``) A Go string.
    :param address: (``string``) A Go string.
    :param username: (``string``) A Go string.
    :param password: (``string``) A Go string.
    :param port: (``int``) An integer.
    :return: The connected service provider.
    :rtype: :go:struct:`CServiceProvider<ydk/types/CServiceProvider>`

.. function:: DisconnectFromRestconfProvider(provider CServiceProvider)

    DisconnectFromRestconfProvider disconnects from RESTCONF device and frees the given service provider

    :param: provider: (:go:struct:`CServiceProvider<ydk/cgopath/CServiceProvider>`) A service provider instance.

.. function:: InitCodecServiceProvider(state *State, entity Entity, repo Repository)

    InitCodecServiceProvider initializes CodecServiceProvider
    
    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param entity: :go:struct:`Entity<ydk/types/Entity>`
    :param repo: (:go:struct:`Repository<ydk/types/Repository>`).
    :return: The root schema node parsed from repository
    :rtype: :go:struct:`RootSchemaNode<ydk/types/RootSchemaNode>`

.. function:: CodecServiceEncode(state *State, entity Entity, rootSchema RootSchemaNode, encoding EncodingFormat)

    CodecServiceEncode encodes entity to XML/JSON payloads based on encoding format passed in

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param entity: (:go:struct:`Entity<ydk/types/Entity>`).
    :param rootSchema: (:go:struct:`RootSchemaNode<ydk/types/RootSchemaNode>`).
    :param encoding: (:go:struct:`EncodingFormat<ydk/types/EncodingFormat>`).
    :return: The resulting payload.
    :rtype: (``string``) A Go string.

.. function:: CodecServiceDecode(state *State, rootSchema RootSchemaNode, payload string, encoding EncodingFormat, topEntity Entity)

    CodecServiceDecode decodes XML/JSON payloads passed in to entity.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param rootSchema: (:go:struct:`RootSchemaNode<ydk/types/RootSchemaNode>`).
    :param payload: (``string``) A Go string.
    :param encoding: (:go:struct:`EncodingFormat<ydk/types/EncodingFormat>`).
    :param topEntity: (:go:struct:`Entity<ydk/types/Entity>`)
    :return: The top level entity from resulting data node.
    :rtype: :go:struct:`Entity<ydk/types/Entity>`

.. function:: ConnectToOpenDaylightProvider(state *State, path, address, username, password string, port int, encoding EncodingFormat, protocol Protocol)

    ConnectToOpenDaylightProvider connects to OpenDaylight device.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param path: (``string``) A Go string.
    :param address: (``string``) A Go string.
    :param username: (``string``) A Go string.
    :param password: (``string``) A Go string.
    :param port: (``int``) An integer.
    :param encoding: (:go:struct:`EncodingFormat<ydk/types/EncodingFormat>`).
    :param protocol: (:go:struct:`Protocol<ydk/types/Protocol>`).
    :return: The connected service provider.
    :rtype: :go:struct:`COpenDaylightServiceProvider<ydk/types/COpenDaylightServiceProvider>`

.. function:: DisconnectFromOpenDaylightProvider(provider COpenDaylightServiceProvider)

    DisconnectFromOpenDaylightProvider disconnects from OpenDaylight device and frees allocated memory.

    :param provider: (:go:struct:`COpenDaylightServiceProvider<ydk/types/COpenDaylightServiceProvider>`).

.. function:: OpenDaylightServiceProviderGetNodeIDs(state *State, provider COpenDaylightServiceProvider)

    A getter function for the node ids given the opendaylight service provider.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param provider: (:go:struct:`COpenDaylightServiceProvider<ydk/types/COpenDaylightServiceProvider>`).
    :returns: A slice of Go strings representing node ids.
    :rtype: ``[]string``

.. function:: OpenDaylightServiceProviderGetNodeProvider(provider COpenDaylightServiceProvider, nodeID string)

    A getter function for the node provider given the opendaylight service provider and node id.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution
    :param provider: (:go:struct:`COpenDaylightServiceProvider<ydk/types/COpenDaylightServiceProvider>`).
    :param nodeID: (``string``) A Go string.
    :return: The service provider.
    :rtype: :go:struct:`CServiceProvider<ydk/types/CServiceProvider>`

.. function:: AddCState(state *State)

    AddCState creates and adds cstate to given state.

    :param state: (pointer to :go:struct:`State<ydk/types/State>`) Current state of execution